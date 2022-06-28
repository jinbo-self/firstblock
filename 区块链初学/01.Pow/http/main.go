package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//设置难度系数
const difficulty = 4

//定义区块

type Block struct {
	//区块高度

	Index int

	Timestanp string

	//当前交易信息
	BMP int

	//当前哈希
	HashCode string

	//上一个哈希
	PreHash string

	//前导0
	Diff int

	//随机值
	Nonce int
}

//数组存储，，一般用链表.
var Blockchain []Block

//加锁，处理并发
var mutex = &sync.Mutex{}

//生成区块
func generateBlock(oldBlock Block, BMP int) Block {
	var newBlock Block
	newBlock.PreHash = oldBlock.HashCode
	newBlock.Timestanp = time.Now().String()

	newBlock.Index = oldBlock.Index + 1
	newBlock.BMP = BMP

	newBlock.Diff = difficulty

	for i := 0; ; i++ {

		hash := calculateHash(newBlock)
		fmt.Println(hash)
		if isHashValid(hash, newBlock.Diff) {
			fmt.Println("挖矿成功")

			newBlock.HashCode = hash

			//
			return newBlock
		}

		//每挖一次。随机值加1
		newBlock.Nonce++
	}
}

//生成哈希值
func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestanp +
		strconv.Itoa(block.Nonce) + strconv.Itoa(block.BMP) + block.PreHash

	sha := sha256.New()
	sha.Write([]byte(record))

	hashed := sha.Sum(nil)

	return hex.EncodeToString(hashed)
}

//判断哈希值的前导0个数和难度系数是否一致
func isHashValid(hash string, diff int) bool {
	prefix := strings.Repeat("0", diff)
	return strings.HasPrefix(hash, prefix)
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
	go func() {
		//创世区块
		genesisBlock := Block{}
		genesisBlock = Block{
			0, time.Now().String(), 0, calculateHash(genesisBlock), "", difficulty, 0,
		}
		mutex.Lock()
		Blockchain = append(Blockchain, genesisBlock)
		mutex.Unlock()
		//格式化输出到控制台
		spew.Dump(genesisBlock)

	}()
	run()

}

//http启动函数
func run() error {
	//处理get或者post请求的回调
	mux := makeMuxRouter()
	httpAddr := os.Getenv("ADDR")
	log.Println("Listening on", os.Getenv("ADDR"))
	ser := &http.Server{
		Addr:         ":" + httpAddr,
		Handler:      mux, //回调函数的句柄
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,

		//最大响应头,1MB
		MaxHeaderBytes: 1 << 20,
	}

	//监听服务
	if err := ser.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

//回调函数
func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()

	//get读，post写
	muxRouter.HandleFunc("/", handGetBlockchain).Methods("GET")

	muxRouter.HandleFunc("/", handWriteBlock).Methods("POST")

	return muxRouter
}

//处理区块链的信息，查询区块,get
func handGetBlockchain(w http.ResponseWriter, r *http.Request) {
	//转json

	bytes, err := json.MarshalIndent(Blockchain, "", "\t")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

//声明post发送的数据类型
type MEssage struct {
	BMP int
}

//处理http的post请求
func handWriteBlock(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var message MEssage

	//从request中读取json数据,创建解码器，匹配自定义类型
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&message); err != nil {
		//返回错误
		respondWithJSON(writer, request, http.StatusNotFound, request.Body)
		return

	}

	//释放资源
	defer request.Body.Close()

	//锁
	mutex.Lock()

	//创建新区块
	newBlock := generateBlock(Blockchain[len(Blockchain)-1], message.BMP)
	mutex.Unlock()

	//判断区块合法性
	if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
		//将区块真正添加到链上
		Blockchain = append(Blockchain, newBlock)
		//格式化输出
		spew.Dump(Blockchain)
	}
	//返回响应信息
	respondWithJSON(writer, request, http.StatusCreated, newBlock)

}

//若错误，服务器返回500
func respondWithJSON(writer http.ResponseWriter, request *http.Request, code int, inter interface{}) {
	//设置响应头
	writer.Header().Set("Content-Type", "application/json")
	//格式化输出json
	response, err := json.MarshalIndent(inter, "", "\t")

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("http 500:server error"))
		return
	}

	//返回指定错误码
	writer.WriteHeader(code)

	//返回指定数据
	writer.Write(response)
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.HashCode != newBlock.PreHash {
		return false
	}

	//或者再次计算哈希值
	if calculateHash(newBlock) != newBlock.HashCode {
		return false
	}
	return true
}

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//申明一个区块

type Block struct {
	//上一个区块的哈希
	PreHash string

	HashCode string

	TimeStanp string

	//当前网络难度系数
	//控制哈希值有几个前导0,计算哈希值要一样
	Diff int

	//交易信息
	Data string

	//高度
	Index int

	//随机值，让下个区块的哈希不要相同
	Nonce int
}

//创建创世区块（第一个区块）
func GenerateFirstBlock(data string) Block {
	var firstblock Block
	firstblock.PreHash = "0"
	firstblock.TimeStanp = time.Now().String()
	firstblock.Diff = 4

	firstblock.Data = data
	firstblock.Index = 1
	firstblock.Nonce = 0

	//当前块的hash
	//用sha256算一个真正的hash
	firstblock.HashCode = GenerationHashValue(firstblock)
	return firstblock
}

func GenerationHashValue(bloc Block) string {
	var hashdata = strconv.Itoa(bloc.Index) + strconv.Itoa(bloc.Nonce) +
		strconv.Itoa(bloc.Diff) + bloc.TimeStanp

	//哈希算法
	var sha = sha256.New()
	sha.Write([]byte(hashdata))
	hashed := sha.Sum(nil)

	return hex.EncodeToString(hashed)
}
func main() {
	//创建创世区块
	block := GenerateFirstBlock("创世区块")
	fmt.Println(block)
	fmt.Println(block.Data)
	//产生第二个区块
	GenerateNextBlock("第二个区块", block)

}
func GenerateNextBlock(data string, oldBlock Block) Block {
	//产生一个新的区块
	var newBlock Block
	newBlock.TimeStanp = time.Now().String()
	newBlock.Diff = 4

	newBlock.Index = 2
	newBlock.Data = data

	newBlock.PreHash = oldBlock.HashCode

	//一般矿工调整
	newBlock.Nonce = 0
	//利用Pow进行挖矿
	newBlock.HashCode = pow(newBlock.Diff, &oldBlock)
	return newBlock
}

//pow工作量证明算法进行哈希碰撞
func pow(diff int, block *Block) string {
	//不停的去挖矿
	for {
		hash := GenerationHashValue(*block)
		//每挖一次，打印一次哈希值
		fmt.Println(hash)
		if strings.HasPrefix(hash, strings.Repeat("0", diff)) {
			fmt.Println("挖矿成功")
			return hash
		} else {
			block.Nonce++
		}
	}
}

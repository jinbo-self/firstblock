package main

import (
	"区块链初学/01.Pow/Block"
	"区块链初学/01.Pow/Blockchain"
)

func main() {
	first := Block.GenerateFirstBlock("创世区块")

	//生成下一个区块
	second := Block.GenerateNextBlock("第二个区块", first)

	//创建链表
	header := Blockchain.CreateHeadeNode(&first)

	Blockchain.AddNode(&second, header)
	Blockchain.ShowNodes(header)

}

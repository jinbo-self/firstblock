package Blockchain

import (
	"fmt"
	"区块链初学/01.Pow/Block"
)

//通过链表的形式，维护区块链中的业务
type Node struct {

	//指针域
	NextNode *Node
	//数据域
	Data *Block.Block
}

//创建头节点
func CreateHeadeNode(data *Block.Block) *Node {
	headerNdoe := new(Node)
	headerNdoe.NextNode = nil
	//传入data
	headerNdoe.Data = data
	//返回头节点
	return headerNdoe
}

//添加节点
//挖矿成功，添加区块
func AddNode(data *Block.Block, node *Node) *Node {
	var newNode *Node = new(Node)
	newNode.NextNode = nil
	newNode.Data = data

	node.NextNode = newNode

	return newNode
}

func ShowNodes(node *Node) {
	n := node
	for {
		//没有下个节点结束
		if n.NextNode == nil {
			fmt.Println(n.Data)
			break
		} else {
			fmt.Println(n.Data)
			n = n.NextNode
		}
	}
}

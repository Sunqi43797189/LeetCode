package main

import (
	"fmt"
	"leet_code_go/helper"
)

func ExchangeChildNode(node *helper.BinaryTreeNode) *helper.BinaryTreeNode {
	if node == nil {
		return nil
	}
	if node.Left == nil && node.Right == nil {
		return node
	}
	temp := node.Left
	node.Left = node.Right
	node.Right = temp
	ExchangeChildNode(node.Left)
	ExchangeChildNode(node.Right)
	return node
}

func main() {
	root := helper.CommonBinaryTree()
	fmt.Println(root.Left, root.Right)
	ExchangeChildNode(root)
	fmt.Println(root.Left, root.Right)
}

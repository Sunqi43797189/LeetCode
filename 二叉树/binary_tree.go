package main

import (
	"fmt"
	"leet_code_go/helper"
)

func PreOrderRecur(node *helper.BinaryTreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Value)
	PreOrderRecur(node.Left)
	PreOrderRecur(node.Right)
}

func PreOrderUnRecur(node *helper.BinaryTreeNode) {
	if node == nil {
		return
	}
	var arr = []*helper.BinaryTreeNode{node}
	for len(arr) != 0 {
		node := arr[len(arr)-1]
		arr = arr[:len(arr)-1]
		fmt.Println(node.Value)
		if node.Right != nil {
			arr = append(arr, node.Right)
		}
		if node.Left != nil {
			arr = append(arr, node.Left)
		}
	}
}

func InOrderRecur(node *helper.BinaryTreeNode) {
	if node == nil {
		return
	}
	InOrderRecur(node.Left)
	fmt.Println(node.Value)
	InOrderRecur(node.Right)
}

func InOrderUnRecur(node *helper.BinaryTreeNode) {
	if node == nil {
		return
	}
	var arr []*helper.BinaryTreeNode
	for {
		if len(arr) != 0 || node != nil {
			if node != nil {
				arr = append(arr, node)
				node = node.Left
			} else {
				node = arr[len(arr)-1]
				arr = arr[:len(arr)-1]
				fmt.Println(node.Value)
				node = node.Right
			}
		}
	}
}

func PosOrderRecur(node *helper.BinaryTreeNode) {
	if node == nil {
		return
	}
	PosOrderRecur(node.Left)
	PosOrderRecur(node.Right)
	fmt.Println(node.Value)
}

func PosOrderUnRecur(node *helper.BinaryTreeNode) {
	if node == nil {
		return
	}
	var arr = []*helper.BinaryTreeNode{node}
	var arr2 []*helper.BinaryTreeNode

	for len(arr) != 0 {
		node = arr[len(arr)-1]
		arr = arr[:len(arr)-1]
		arr2 = append(arr2, node)
		if node.Left != nil {
			arr = append(arr, node.Left)
		}
		if node.Right != nil {
			arr = append(arr, node.Right)
		}
	}
	for i := len(arr2) - 1; i >= 0; i-- {
		fmt.Println(arr2[i])
	}
}

func main() {
	root := helper.FullBinaryTree()
	fmt.Println(root.Value)
}

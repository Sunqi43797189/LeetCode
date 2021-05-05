package main

import (
	"leet_code_go/helper"
	"math"
)

var (
	preValue = math.MinInt64
)

func CheckBst(node *helper.BinaryTreeNode) bool {
	if node == nil {
		return true
	}
	isLeftBst := CheckBst(node.Left)
	if !isLeftBst {
		return false
	}
	if node.Value <= preValue {
		return false
	} else {
		preValue = node.Value
	}
	return CheckBst(node.Right)

}

// func main() {
// 	root := helper.BinarySearchTree()
// 	result := CheckBst(root)
// 	fmt.Println(result)
// }

package main

import (
	"fmt"
	"leet_code_go/helper"
)

func CheckCBT(root *helper.BinaryTreeNode) bool {
	if root == nil {
		return false
	}
	var leaf bool // 遇到叶子节点就会标记为true
	arr := []*helper.BinaryTreeNode{root}
	for len(arr) != 0 {
		node := arr[0]
		arr = arr[1:]
		if node.Right != nil && node.Left == nil { // 完全二叉树一层不满的情况是节点从左到右排列
			return false
		}
		if leaf && (node.Left != nil || node.Right != nil) { // 遇到了叶子节点后边所有节点都是叶子节点
			return false
		}
		if node.Left != nil {
			arr = append(arr, node.Left)
		}
		if node.Right != nil {
			arr = append(arr, node.Right)
		}
		if node.Left == nil && node.Right == nil {
			leaf = true
		}
	}
	return true
}

func main() {
	root := helper.CompleteBinaryTree()
	result := CheckCBT(root)
	fmt.Println(result)
}

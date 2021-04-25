package main

import (
	"fmt"
	"leet_code_go/helper"
	"math"
)

func BlanceBinaryTree(node *helper.BinaryTreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}
	leftBlanced, lHeight := BlanceBinaryTree(node.Left)
	rightBlanced, rHeight := BlanceBinaryTree(node.Right)
	if leftBlanced && rightBlanced && math.Abs(float64(lHeight-rHeight)) < 2 {
		return true, int(math.Max(float64(lHeight), float64(rHeight))) + 1
	} else {
		return false, 0
	}
}

func main() {
	root := helper.CompleteBinaryTree()
	isBlanced, _ := BlanceBinaryTree(root)
	fmt.Println(isBlanced)
}

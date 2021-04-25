package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	DfsPathSum(root, targetSum - root.Val, &res, []int{root.Val})
	return res
}

func DfsPathSum(node *TreeNode, sum int, res *[][]int, tempArr []int){
	if node.Left == nil && node.Right == nil {
		if sum == 0 {
			a := make([]int, len(tempArr))
			copy(a, tempArr)
			*res = append(*res, a)
		}
		return
	}

	if node.Left != nil {
		tempArr = append(tempArr, node.Left.Val)
		DfsPathSum(node.Left, sum - node.Left.Val, res, tempArr)
		tempArr = tempArr[:len(tempArr) - 1]
	}
	if node.Right != nil {
		tempArr = append(tempArr, node.Right.Val)
		DfsPathSum(node.Right, sum - node.Right.Val, res, tempArr)
		tempArr = tempArr[:len(tempArr) - 1]
	}
}

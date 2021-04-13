package main

import (
	"fmt"
	"leet_code_go/helper"
	"math"
)

func BinaryTreeMaxCountLevelWithMap(node *helper.BinaryTreeNode) (level int, count int) {
	if node == nil {
		return
	}
	queue := []*helper.BinaryTreeNode{node}
	nodeMap := map[*helper.BinaryTreeNode]int{
		node: 1,
	}
	curLevel, curLevelNodes, maxCount := 1, 0, -1
	for ; len(queue) != 0; {
		node := queue[0]
		queue = queue[1:]
		if curLevel == nodeMap[node] {
			curLevelNodes += 1
		} else {
			maxCount = int(math.Max(float64(maxCount), float64(curLevelNodes)))
			curLevel += 1
			curLevelNodes = 1
		}
		if node.Left != nil {
			nodeMap[node.Left] = nodeMap[node] + 1
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			nodeMap[node.Right] = nodeMap[node] + 1
			queue = append(queue, node.Right)
		}
	}
	return curLevel - 1, maxCount
}

func BinaryTreeMaxCountLevel(root *helper.BinaryTreeNode) (count int) {
	if root == nil {
		return
	}
	queue := []*helper.BinaryTreeNode{root}
	var curEnd, nextEnd *helper.BinaryTreeNode
	curEnd, nextEnd, curLevelNodes, maxCount := root, nil, 0, -1
	for ; len(queue) != 0; {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
			nextEnd = node.Left
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			nextEnd = node.Right
		}
		if curEnd == node {
			curEnd = nextEnd
			maxCount = int(math.Max(float64(maxCount), float64(curLevelNodes)))
			curLevelNodes = 1
		} else {
			curLevelNodes++
		}

	}
	return maxCount
}

func main() {
	root := helper.CommonBinaryTree()
	fmt.Println(BinaryTreeMaxCountLevelWithMap(root))
	fmt.Println(BinaryTreeMaxCountLevel(root))
}

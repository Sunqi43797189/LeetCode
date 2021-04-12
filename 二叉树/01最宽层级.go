package main

import (
	"fmt"
	"leet_code_go/helper"
	"math"
)

func BinaryTreeMaxCountLevel(node *helper.BinaryTreeNode) (level int, count int) {
	if node == nil {
		return
	}
	queue := []*helper.BinaryTreeNode{node}
	nodeMap := map[*helper.BinaryTreeNode]int{
		node: 1,
	}
	curLevel, curLevelNodes, maxCount := 1, 0, -1
	for len(queue) != 0 {
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

func main() {
	root := helper.FullBinaryTree()
	fmt.Println(BinaryTreeMaxCountLevel(root))
}

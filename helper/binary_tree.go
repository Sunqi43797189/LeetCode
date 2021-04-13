package helper

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// 满二叉树
func CommonBinaryTree() *BinaryTreeNode {
	node1 := &BinaryTreeNode{Value: 1}
	node2 := &BinaryTreeNode{Value: 2}
	node3 := &BinaryTreeNode{Value: 3}
	node4 := &BinaryTreeNode{Value: 4}
	node5 := &BinaryTreeNode{Value: 5}
	node6 := &BinaryTreeNode{Value: 6}
	node7 := &BinaryTreeNode{Value: 7}
	node8 := &BinaryTreeNode{Value: 8}

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7

	node4.Right = node8

	return node1
}

// 二叉搜索树
func BinarySearchTree() *BinaryTreeNode {
	node1 := &BinaryTreeNode{Value: 1}
	node2 := &BinaryTreeNode{Value: 2}
	node3 := &BinaryTreeNode{Value: 3}
	node4 := &BinaryTreeNode{Value: 4}
	node5 := &BinaryTreeNode{Value: 5}
	node6 := &BinaryTreeNode{Value: 6}
	node7 := &BinaryTreeNode{Value: 7}
	node8 := &BinaryTreeNode{Value: 8}

	node5.Left = node3
	node5.Right = node7

	node7.Left = node6
	node7.Right = node8

	node3.Left = node1
	node3.Right = node4

	node1.Right = node2
	return node5
}

// 满二叉树
func CompleteBinaryTree()  *BinaryTreeNode {
	node1 := &BinaryTreeNode{Value: 1}
	node2 := &BinaryTreeNode{Value: 2}
	node3 := &BinaryTreeNode{Value: 3}
	node4 := &BinaryTreeNode{Value: 4}
	node5 := &BinaryTreeNode{Value: 5}
	node6 := &BinaryTreeNode{Value: 6}
	node7 := &BinaryTreeNode{Value: 7}
	node8 := &BinaryTreeNode{Value: 8}

	node1.Left = node2
	node1.Right = node3

	node2.Left = node4
	node2.Right = node5

	node3.Left = node6
	node3.Right = node7

	node4.Left = node8
	return node1
}

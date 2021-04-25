package helper

type LinkListNode struct {
	Value int
	Next  *LinkListNode
}

func NoCircleLinkList() *LinkListNode {
	node1 := LinkListNode{Value: 1}
	node2 := LinkListNode{Value: 2}
	node3 := LinkListNode{Value: 3}
	node4 := LinkListNode{Value: 4}
	node5 := LinkListNode{Value: 5}
	node6 := LinkListNode{Value: 6}
	node7 := LinkListNode{Value: 7}
	node8 := LinkListNode{Value: 8}
	node9 := LinkListNode{Value: 9}

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6
	node6.Next = &node7
	node7.Next = &node8
	node8.Next = &node9

	return &node1
}

func CircleLinkList() *LinkListNode {
	node1 := LinkListNode{Value: 1}
	node2 := LinkListNode{Value: 2}
	node3 := LinkListNode{Value: 3}
	node4 := LinkListNode{Value: 4}
	node5 := LinkListNode{Value: 5}
	node6 := LinkListNode{Value: 6}
	node7 := LinkListNode{Value: 7}
	node8 := LinkListNode{Value: 8}
	node9 := LinkListNode{Value: 9}

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6
	node6.Next = &node7
	node7.Next = &node8
	node8.Next = &node9
	node9.Next = &node6

	return &node1
}

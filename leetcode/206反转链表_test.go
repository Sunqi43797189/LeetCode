package leetcode

import (
	"fmt"
	"testing"
)

func Test_reverseList(t *testing.T) {
	node5 := ListNode{Val: 5, Next: nil}
	node4 := ListNode{Val: 4, Next: &node5}
	node3 := ListNode{Val: 3, Next: &node4}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}

	PrintListNode(&node1)
	fmt.Println("**************************")
	resultNode := reverseList(&node1)
	fmt.Println("**************************")
	PrintListNode(resultNode)
}

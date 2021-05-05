package leetcode

import (
	"fmt"
)

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

/*
https://leetcode-cn.com/problems/reverse-linked-list/
1.定义一个节点prev，用于保存上一个节点
2.先保存当前节点的下一个节点temp，便于后续调整
3. 当前节点的下一个节点指向prev
4. 然后将当前节点作为上一个节点赋值到prev
5. 然后 将刚刚保存的节点作为当前节点，循环继续
6. 当前节点为空也就是上一次循环保存的下一个节点为nil，然后终止循环
7. 此时 prev就是头结点了，return
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}

func PrintListNode(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

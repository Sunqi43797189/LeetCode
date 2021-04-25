package link_list

import (
	"leet_code_go/helper"
)

func IsLinkListCircle(head *helper.LinkListNode) (bool, *helper.LinkListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false, nil
	}

	fast := head.Next.Next
	slow := head.Next

	for fast != slow {
		if fast.Next == nil || fast.Next.Next == nil {
			return false, nil
		}
		fast = fast.Next.Next
		slow = slow.Next
	} // 两个走到同一个节点时，快指针回到头部，然后一起走，直到两个指针再次重合
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return true, slow
}

package Q135

import "reflect"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	tmp := new(ListNode)
	tmp = head
	re := reverse(tmp)

	return reflect.DeepEqual(re, head)
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}

	return prev
}

package Q135

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {

	for {

	}

	return false
}

func visit(head *ListNode) int {
	if head.Next != nil {
		visit(head.Next)
	}

	return head.Val
}

// func reverse(head *ListNode) *ListNode {
// 	var prev *ListNode

// 	for head != nil {
// 		head.Next, prev, head = prev, head, head.Next
// 	}

// 	return prev
// }

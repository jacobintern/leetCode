package Q122

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	return traversalNode(head, nil)
}

func traversalNode(head, res *ListNode) *ListNode {
	if head == nil {
		return res
	}

	next := head.Next
	head.Next = res
	res = head

	return traversalNode(next, res)
}

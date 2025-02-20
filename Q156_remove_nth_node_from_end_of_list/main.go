package Q156

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	remove(&head, &n)
	return head
}

func remove(node **ListNode, n *int) {
	if *node == nil {
		return
	}

	remove(&(*node).Next, n)

	if *n--; *n == 0 {
		*node = (*node).Next
	}
}

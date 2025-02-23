package Q158

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	l := len(lists)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return lists[0]
	}
	mid := l / 2
	return merge(mergeKLists(lists[0:mid]), mergeKLists(lists[mid:]))
}

func merge(n1, n2 *ListNode) *ListNode {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}

	if n1.Val > n2.Val {
		n1.Next = merge(n1.Next, n2)
		return n1
	}

	n2.Next = merge(n1, n2.Next)
	return n2
}

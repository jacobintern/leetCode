package Q133

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head, point := head.Next, head

	for point != nil && point.Next != nil {
		tmp := new(ListNode)
		tmp.Val = point.Val
		tmp.Next = point.Next.Next
		point = point.Next
		point.Next = tmp

		point = point.Next.Next
	}

	return head
}

func swapPairs3(head *ListNode) *ListNode {
	head, point := head.Next, head
	var prev *ListNode

	for point != nil && point.Next != nil {
		tmp := point.Next

		point.Next, tmp.Next = tmp.Next, point

		if prev != nil {
			prev.Next = tmp
		}

		prev, point = point, point.Next
	}

	return head
}

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	point := head

	for point != nil {
		if point.Next != nil {
			tmp := point.Val
			point.Val = point.Next.Val
			point.Next.Val = tmp
			point = point.Next
		}

		if point != nil {
			point = point.Next
		}
	}

	return head
}

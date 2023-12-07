package Q134

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	res, ok := hasCycle(head)
	if !ok {
		return nil
	}
	l1, l2, tmp := 0, 1, res
	// 找出長度
	for {
		l1++
		tmp = tmp.Next
		if res == tmp {
			break
		}
	}

	if head == res {
		return res
	}
	p := head
	for {
		l2++
		p = p.Next
		if res == p {
			break
		}
	}

	// 結果就是總長剪去迴圈長度
	count := l2 - l1 + 1
	for {
		if count <= 0 {
			break
		}
		head = head.Next
		count--
	}

	return head
}

func detectCycle2(head *ListNode) *ListNode {
	res, ok := hasCycle(head)
	if !ok {
		return nil
	}

	for head != res {
		head = head.Next
		res = res.Next
	}
	return head
}

func hasCycle(head *ListNode) (*ListNode, bool) {
	p1, p2 := head, head

	for p1 != nil && p1.Next != nil {
		p1, p2 = p1.Next, p2.Next.Next
		if p1 == p2 {
			return p1, true
		}
	}

	return nil, false
}

package Q19

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var r []int
	r = append(r, head.Val)
	i := 0
	for {
		if r[i] != head.Val {
			r = append(r, head.Val)
			i++
		}
		if head.Next == nil {
			break
		}
		head = head.Next
	}
	node := &ListNode{
		Val: r[0],
	}

	temp := node

	for i := 1; i < len(r); i++ {
		temp.Next = &ListNode{
			Val: r[i],
		}
		temp = temp.Next
	}

	return node
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	Next := head.Next
	for Next != nil {
		if Next.Val == head.Val {
			Next = Next.Next
		} else {
			break
		}
	}
	head.Next = deleteDuplicates(Next)
	return head
}

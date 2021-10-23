package Q7

type ListNode struct {
	Val  int
	Next *ListNode
}

var arr []int

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	resArr := OutputArray(l1, l2)

	r := &ListNode{}
	tmp := r
	for i, v := range resArr {
		if i == 0 {
			r.Val = v
		} else {
			tmp.Next = &ListNode{Val: v}
			tmp = tmp.Next
		}
	}
	arr = nil
	return r
}

func OutputArray(l1 *ListNode, l2 *ListNode) []int {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l2 == nil {
		arr = append(arr, l1.Val)
		l1 = l1.Next
	} else if l1 == nil {
		arr = append(arr, l2.Val)
		l2 = l2.Next
	} else if l1.Val < l2.Val {
		arr = append(arr, l1.Val)
		l1 = l1.Next
	} else if l1.Val > l2.Val {
		arr = append(arr, l2.Val)
		l2 = l2.Next
	} else {
		arr = append(arr, l1.Val)
		l1 = l1.Next
	}
	OutputArray(l1, l2)
	return arr
}

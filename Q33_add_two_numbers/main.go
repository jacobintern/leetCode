package Q33

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	cur := result
	sum := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: sum % 10}
		if sum > 9 {
			sum = 1
		} else {
			sum = 0
		}
		cur = cur.Next
	}
	if sum > 0 {
		result.Val = sum
	}
	return result.Next
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	result := &ListNode{}
// 	for l1 != nil || l2 != nil {
// 		sum := 0
// 		if l1 != nil {
// 			sum += l1.Val
// 			l1 = l1.Next
// 		}
// 		if l2 != nil {
// 			sum += l2.Val
// 			l2 = l2.Next
// 		}
// 		if result.Next != nil {
// 			result.Next.Val += sum % 10
// 		} else {
// 			result.Next = &ListNode{Val: sum % 10}
// 		}
// 		result = result.Next
// 		if result.Val > 9 {
// 			result.Next = &ListNode{Val: 1}
// 		}
// 		fmt.Println(result.Val)

// 	}

// 	return result
// }

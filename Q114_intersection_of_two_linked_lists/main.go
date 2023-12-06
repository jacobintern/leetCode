package Q114

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil && headB == nil {
		return nil
	}

	if headA.Next != nil && headB.Next != nil {
		GetIntersectionNode(headA.Next, headB.Next)
	} else if headA.Next == nil && headB.Next != nil {
		GetIntersectionNode(headA, headB.Next)
	} else if headA.Next != nil && headB.Next == nil {
		GetIntersectionNode(headA.Next, headB)
	}

	fmt.Println("A Node:", headA.Val)
	fmt.Println("B Node:", headB.Val)

	return nil
}

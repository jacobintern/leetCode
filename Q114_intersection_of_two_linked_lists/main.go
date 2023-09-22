package Q114

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA != nil {
		fmt.Println("A Node:", headA.Val)
	}
	if headB != nil {
		fmt.Println("B Node:", headB.Val)
	}

	if headA == nil && headB == nil {
		return nil
	}

	GetIntersectionNode(headA.Next, headB.Next)

	return nil
}

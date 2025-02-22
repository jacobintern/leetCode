package Q114

import (
	"reflect"
	"testing"
)

type parameters struct {
	headA *ListNode
	headB *ListNode
}

func Test1(t *testing.T) {

	params := parameters{headA: &ListNode{Val: 4,
		Next: &ListNode{Val: 1,
			Next: &ListNode{Val: 8,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 5}}}}},
		headB: &ListNode{Val: 5,
			Next: &ListNode{Val: 6,
				Next: &ListNode{Val: 1,
					Next: &ListNode{Val: 8,
						Next: &ListNode{Val: 4,
							Next: &ListNode{Val: 5}}}}}}}

	expected := &ListNode{Val: 8,
		Next: &ListNode{Val: 4,
			Next: &ListNode{Val: 5}}}

	if testResult := GetIntersectionNode(params.headA, params.headB); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

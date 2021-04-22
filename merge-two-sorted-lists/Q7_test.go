package Q7

import (
	"reflect"
	"testing"
)

type parameters struct {
	arr1 ListNode
	arr2 ListNode
}

func Test1(t *testing.T) {
	params := parameters{
		arr1: ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
		arr2: ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
	}
	expected := ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}}

	if reflect.DeepEqual(expected, mergeTwoLists(&params.arr1, &params.arr2)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

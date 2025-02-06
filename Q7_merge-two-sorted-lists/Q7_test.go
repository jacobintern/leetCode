package Q7

import (
	"reflect"
	"testing"
)

type parameters struct {
	node1 ListNode
	node2 ListNode
}

func Test1(t *testing.T) {
	params := parameters{
		node1: ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
		node2: ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
	}
	expected := ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}}

	if reflect.DeepEqual(expected, mergeTwoLists(&params.node1, &params.node2)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

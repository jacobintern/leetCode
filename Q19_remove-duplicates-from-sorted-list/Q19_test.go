package Q19

import (
	"reflect"
	"testing"
)

type parameters struct {
	head ListNode
}

func Test1(t *testing.T) {
	params := parameters{
		head: ListNode{Val: 1,
			Next: &ListNode{Val: 1,
				Next: &ListNode{Val: 2},
			},
		},
	}
	expected := &ListNode{Val: 1, Next: &ListNode{Val: 2}}

	if reflect.DeepEqual(expected, deleteDuplicates(&params.head)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		head: ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}},
	}
	expected := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}

	if reflect.DeepEqual(expected, deleteDuplicates(&params.head)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

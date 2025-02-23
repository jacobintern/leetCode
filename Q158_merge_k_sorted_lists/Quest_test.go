package Q158

import (
	"reflect"
	"testing"
)

type parameters struct {
	lists []*ListNode
}

func Test1(t *testing.T) {

	params := parameters{lists: []*ListNode{
		{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		{Val: 2, Next: &ListNode{Val: 6}},
	}}

	expected := &ListNode{Val: 1,
		Next: &ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 4,
							Next: &ListNode{Val: 5,
								Next: &ListNode{Val: 6}}}}}}}}

	if testResult := mergeKLists(params.lists); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{lists: []*ListNode{}}

	expected := &ListNode{}

	if testResult := mergeKLists(params.lists); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{lists: []*ListNode{{}}}

	expected := &ListNode{}

	if testResult := mergeKLists(params.lists); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

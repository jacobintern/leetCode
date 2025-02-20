package Q155

import (
	"reflect"
	"testing"
)

type parameters struct {
	head *ListNode
	n    int
}

func Test1(t *testing.T) {

	params := parameters{
		head: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		},
		n: 2,
	}

	expected := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
	}

	if testResult := removeNthFromEnd(params.head, params.n); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{
		head: &ListNode{Val: 1},
		n:    1,
	}

	expected := &ListNode{}

	if testResult := removeNthFromEnd(params.head, params.n); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{
		head: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		},
		n: 1,
	}

	expected := &ListNode{Val: 1}

	if testResult := removeNthFromEnd(params.head, params.n); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

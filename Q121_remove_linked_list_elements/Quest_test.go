package Q121

import (
	"reflect"
	"testing"
)

type parameters struct {
	head *ListNode
	val  int
}

func Test1(t *testing.T) {

	params := parameters{head: &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 6,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5,
							Next: &ListNode{Val: 6,
								Next: &ListNode{}}}}}}}},
		val: 6}

	expected := &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 5}}}}}

	if testResult := removeElements(params.head, params.val); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{head: &ListNode{},
		val: 1}

	expected := &ListNode{}

	if testResult := removeElements(params.head, params.val); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}
func Test3(t *testing.T) {

	params := parameters{head: &ListNode{Val: 7,
		Next: &ListNode{Val: 7,
			Next: &ListNode{Val: 7,
				Next: &ListNode{Val: 7,
					Next: &ListNode{}}}}},
		val: 7}

	expected := &ListNode{}

	if testResult := removeElements(params.head, params.val); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{head: &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{}}},
		val: 1}

	expected := &ListNode{Val: 2}

	if testResult := removeElements(params.head, params.val); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

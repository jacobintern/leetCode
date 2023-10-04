package Q122

import (
	"reflect"
	"testing"
)

type parameters struct {
	head *ListNode
}

func Test1(t *testing.T) {

	params := parameters{head: &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 5}}}}}}

	expected := &ListNode{Val: 5,
		Next: &ListNode{Val: 4,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 1}}}}}

	if testResult := reverseList(params.head); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{head: &ListNode{Val: 1,
		Next: &ListNode{Val: 2}}}

	expected := &ListNode{Val: 2,
		Next: &ListNode{Val: 1}}

	if testResult := reverseList(params.head); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{head: &ListNode{}}

	expected := &ListNode{}

	if testResult := reverseList(params.head); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

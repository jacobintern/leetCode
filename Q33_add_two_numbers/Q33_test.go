package Q33

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
		node1: ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
		node2: ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
	}

	expected := &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}}

	if testResult := addTwoNumbers(&params.node1, &params.node2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{
		node1: ListNode{Val: 0},
		node2: ListNode{Val: 0},
	}

	expected := &ListNode{Val: 0}

	if testResult := addTwoNumbers(&params.node1, &params.node2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{
		node1: ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}},
		node2: ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}},
	}

	expected := &ListNode{Val: 0}

	if testResult := addTwoNumbers(&params.node1, &params.node2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

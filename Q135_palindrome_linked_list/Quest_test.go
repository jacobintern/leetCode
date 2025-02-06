package Q135

import (
	"reflect"
	"testing"
)

type parameters struct {
	head *ListNode
}

func Test1(t *testing.T) {

	params := parameters{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}}

	expected := true

	if testResult := isPalindrome(params.head); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{head: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}

	expected := false

	if testResult := isPalindrome(params.head); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

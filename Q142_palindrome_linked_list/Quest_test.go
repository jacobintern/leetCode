package Q142

import (
	"reflect"
	"testing"
)

type parameters struct {
	head *ListNode
}

func Test1(t *testing.T) {

	params := parameters{
		head: &ListNode{Val: 1,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 1},
				},
			},
		}}

	expected := 0

	if testResult := isPalindrome(params.head); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{
		head: &ListNode{Val: 3,
			Next: &ListNode{Val: 2},
		}}

	expected := 0

	if testResult := isPalindrome(params.head); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

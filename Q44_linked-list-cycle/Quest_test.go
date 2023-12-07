package Q44

import (
	"reflect"
	"testing"
)

type parameters struct {
	head ListNode
}

func Test1(t *testing.T) {

	params := parameters{
		head: ListNode{Val: 3,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 0,
					Next: &ListNode{Val: -4},
				},
			},
		},
	}

	expected := true

	if testResult := hasCycle(&params.head); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

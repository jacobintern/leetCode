package Q44

import (
	"reflect"
	"testing"
)

type parameters struct {
	head Listnode
}

func Test1(t *testing.T) {

	params := parameters{
		head: Listnode{Val: 3,
			Next: &Listnode{Val: 2,
				Next: &Listnode{Val: 0,
					Next: &Listnode{Val: -4},
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

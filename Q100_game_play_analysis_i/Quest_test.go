package Q100

import (
	"reflect"
	"testing"
)

type parameters struct {
	root TreeNode
}

func Test1(t *testing.T) {

	params := parameters{root: TreeNode{Val: 1,
		Right: &TreeNode{Val: 2,
			Left: &TreeNode{Val: 2}}}}

	expected := 0

	if testResult := findMode(&params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{root: TreeNode{Val: 0}}

	expected := 0

	if testResult := findMode(&params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

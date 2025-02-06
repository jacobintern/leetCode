package Q138

import (
	"reflect"
	"testing"
)

type parameters struct {
	root   *TreeNode
	target int
}

func Test1(t *testing.T) {

	params := parameters{root: &TreeNode{Val: 5,
		Left:  &TreeNode{},
		Right: &TreeNode{},
	}}

	expected := 0

	if testResult := findTarget(params.root, params.target); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

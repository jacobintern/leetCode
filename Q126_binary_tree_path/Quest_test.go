package Q126

import (
	"reflect"
	"testing"
)

type parameters struct {
	root *TreeNode
}

func Test1(t *testing.T) {

	params := parameters{root: &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3}}}

	expected := []string{"1->2->5", "1->3"}

	if testResult := binaryTreePaths(params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{root: &TreeNode{Val: 1}}

	expected := []string{"1"}

	if testResult := binaryTreePaths(params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

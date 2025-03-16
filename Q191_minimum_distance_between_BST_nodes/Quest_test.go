package Q191

import (
	"reflect"
	"testing"
)

type parameters struct {
	root *TreeNode
}

func Test1(t *testing.T) {

	params := parameters{root: &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 6}}}

	expected := 1

	if testResult := minDiffInBST(params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{root: &TreeNode{Val: 1,
		Left: &TreeNode{Val: 0},
		Right: &TreeNode{Val: 48,
			Left:  &TreeNode{Val: 12},
			Right: &TreeNode{Val: 49}}}}

	expected := 1

	if testResult := minDiffInBST(params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

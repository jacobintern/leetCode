package Q125

import (
	"reflect"
	"testing"
)

type parameters struct {
	root *TreeNode
}

func Test1(t *testing.T) {

	params := parameters{root: &TreeNode{Val: 4,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9}}}}

	expected := &TreeNode{Val: 4,
		Left: &TreeNode{Val: 7,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 6}},
		Right: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 1}}}

	if testResult := invertTree(params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{root: &TreeNode{Val: 2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3}}}

	expected := &TreeNode{Val: 2,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 1}}

	if testResult := invertTree(params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{root: &TreeNode{}}

	expected := &TreeNode{}

	if testResult := invertTree(params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{root: &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2}}}

	expected := &TreeNode{Val: 1,
		Right: &TreeNode{Val: 2}}

	if testResult := invertTree(params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

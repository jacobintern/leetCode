package Q130

import (
	"reflect"
	"testing"
)

type parameters struct {
	root1 *TreeNode
	root2 *TreeNode
}

func Test1(t *testing.T) {

	params := parameters{
		root1: &TreeNode{Val: 1,
			Left: &TreeNode{Val: 3,
				Left: &TreeNode{Val: 5}},
			Right: &TreeNode{Val: 2}},
		root2: &TreeNode{Val: 2,
			Left: &TreeNode{Val: 1,
				Right: &TreeNode{Val: 4}},
			Right: &TreeNode{Val: 3,
				Right: &TreeNode{Val: 7}}}}

	expected := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 4,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 5,
			Right: &TreeNode{Val: 7}}}

	if testResult := mergeTrees(params.root1, params.root2); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{
		root1: &TreeNode{Val: 1},
		root2: &TreeNode{Val: 1,
			Left: &TreeNode{Val: 2}}}

	expected := &TreeNode{Val: 2,
		Left: &TreeNode{Val: 2}}

	if testResult := mergeTrees(params.root1, params.root2); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

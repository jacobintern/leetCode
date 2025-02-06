package Q110

import (
	"reflect"
	"testing"
)

type parameters struct {
	root    TreeNode
	subRoot TreeNode
}

func Test1(t *testing.T) {

	params := parameters{
		root: TreeNode{Val: 3,
			Left: &TreeNode{Val: 4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2}},
			Right: &TreeNode{Val: 5}},
		subRoot: TreeNode{Val: 4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2}}}

	expected := true

	if testResult := isSubtree(&params.root, &params.subRoot); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{
		root: TreeNode{Val: 3,
			Left: &TreeNode{Val: 4,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2,
					Left: &TreeNode{Val: 0}}},
			Right: &TreeNode{Val: 5}},
		subRoot: TreeNode{Val: 4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2}}}

	expected := false

	if testResult := isSubtree(&params.root, &params.subRoot); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

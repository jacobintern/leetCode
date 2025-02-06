package Q118

import (
	"reflect"
	"testing"
)

type parameters struct {
	root TreeNode
}

func Test1(t *testing.T) {

	params := parameters{root: TreeNode{Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}}

	expected := true

	if testResult := isBalanced(&params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{root: TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left: &TreeNode{Val: 3,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 4}},
			Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2},
	}}

	expected := false

	if testResult := isBalanced(&params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}
func Test3(t *testing.T) {

	params := parameters{root: TreeNode{}}

	expected := true

	if testResult := isBalanced(&params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

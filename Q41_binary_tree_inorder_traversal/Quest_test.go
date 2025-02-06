package Q41

import (
	"reflect"
	"testing"
)

type parameters struct {
	tree *TreeNode
}

func Test1(t *testing.T) {

	params := parameters{
		tree: &TreeNode{Val: 1,
			Left: nil,
			Right: &TreeNode{Val: 2,
				Left:  &TreeNode{Val: 3},
				Right: nil,
			},
		},
	}

	expected := []int{1, 3, 2}

	if testResult := inorderTraversal(params.tree); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{
		tree: nil,
	}

	expected := []int{}

	if testResult := inorderTraversal(params.tree); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{
		tree: &TreeNode{Val: 1},
	}

	expected := []int{1}

	if testResult := inorderTraversal(params.tree); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{
		tree: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 2},
			Right: nil,
		},
	}

	expected := []int{2, 1}

	if testResult := inorderTraversal(params.tree); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test5(t *testing.T) {

	params := parameters{
		tree: &TreeNode{Val: 1,
			Left:  nil,
			Right: &TreeNode{Val: 2},
		},
	}

	expected := []int{1, 2}

	if testResult := inorderTraversal(params.tree); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test6(t *testing.T) {

	params := parameters{
		tree: &TreeNode{Val: 4,
			Left: &TreeNode{Val: 2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3}},
			Right: &TreeNode{Val: 6,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 7}},
		},
	}

	expected := []int{1, 2, 3, 4, 5, 6, 7}

	if testResult := inorderTraversal(params.tree); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

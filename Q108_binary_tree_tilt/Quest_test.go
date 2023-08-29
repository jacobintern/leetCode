package Q108

import (
	"reflect"
	"testing"
)

type parameters struct {
	tree TreeNode
}

func Test1(t *testing.T) {

	params := parameters{tree: TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
		Right: &TreeNode{Val: 3, Left: nil, Right: nil},
	}}

	expected := 1

	if testResult := findTilt(&params.tree); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{tree: TreeNode{Val: 4,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 3, Left: &TreeNode{}, Right: &TreeNode{}},
			Right: &TreeNode{Val: 5, Left: &TreeNode{}, Right: &TreeNode{}}},
		Right: &TreeNode{Val: 9,
			Left:  &TreeNode{},
			Right: &TreeNode{Val: 7, Left: &TreeNode{}, Right: &TreeNode{}}},
	}}

	expected := 15

	if testResult := findTilt(&params.tree); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{tree: TreeNode{Val: 21,
		Left: &TreeNode{Val: 7,
			Left: &TreeNode{Val: 1,
				Left:  &TreeNode{Val: 3, Left: &TreeNode{}, Right: &TreeNode{}},
				Right: &TreeNode{Val: 3, Left: &TreeNode{}, Right: &TreeNode{}}},
			Right: &TreeNode{Val: 1, Left: &TreeNode{}, Right: &TreeNode{}}},
		Right: &TreeNode{Val: 14,
			Left:  &TreeNode{Val: 2, Left: &TreeNode{}, Right: &TreeNode{}},
			Right: &TreeNode{Val: 2, Left: &TreeNode{}, Right: &TreeNode{}}},
	}}

	expected := 9

	if testResult := findTilt(&params.tree); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

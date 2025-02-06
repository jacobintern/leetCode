package Q23

import (
	"reflect"
	"testing"
)

type parameters struct {
	tree *TreeNode
}

func Test1(t *testing.T) {
	params := parameters{
		tree: &TreeNode{Val: 3,
			Left: &TreeNode{Val: 9},
			Right: &TreeNode{Val: 20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7}}},
	}
	expected := 3

	if reflect.DeepEqual(expected, maxDepth(params.tree)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		tree: &TreeNode{Val: 1,
			Left:  nil,
			Right: &TreeNode{Val: 2}},
	}
	expected := 2

	if reflect.DeepEqual(expected, maxDepth(params.tree)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}
func Test3(t *testing.T) {
	params := parameters{
		tree: nil,
	}
	expected := 0

	if reflect.DeepEqual(expected, maxDepth(params.tree)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}
func Test4(t *testing.T) {
	params := parameters{
		tree: &TreeNode{Val: 0},
	}
	expected := 1

	if reflect.DeepEqual(expected, maxDepth(params.tree)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

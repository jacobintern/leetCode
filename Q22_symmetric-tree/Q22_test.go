package Q22

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
			Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}},
	}
	expected := true

	if reflect.DeepEqual(expected, isSymmetric(params.tree)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		tree: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
			Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
	}
	expected := false

	if reflect.DeepEqual(expected, isSymmetric(params.tree)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

package Q21

import (
	"reflect"
	"testing"
)

type parameters struct {
	tree1 *TreeNode
	tree2 *TreeNode
}

func Test1(t *testing.T) {
	params := parameters{
		tree1: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
		tree2: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
	}
	expected := true

	if reflect.DeepEqual(expected, isSameTree(params.tree1, params.tree2)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		tree1: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
		tree2: &TreeNode{Val: 1, Right: &TreeNode{Val: 3}},
	}
	expected := false

	if reflect.DeepEqual(expected, isSameTree(params.tree1, params.tree2)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test3(t *testing.T) {
	params := parameters{
		tree1: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
		tree2: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}},
	}
	expected := false

	if reflect.DeepEqual(expected, isSameTree(params.tree1, params.tree2)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

package Q24

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
				Right: &TreeNode{Val: 7},
			},
		},
	}
	expected := 2

	if reflect.DeepEqual(expected, minDepth(params.tree)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		tree: &TreeNode{Val: 2,
			Left: nil,
			Right: &TreeNode{Val: 3,
				Left: nil,
				Right: &TreeNode{Val: 4,
					Left: nil,
					Right: &TreeNode{Val: 5,
						Left:  nil,
						Right: &TreeNode{Val: 6},
					},
				},
			},
		},
	}
	expected := 5

	if reflect.DeepEqual(expected, minDepth(params.tree)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test3(t *testing.T) {
	params := parameters{
		tree: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 2},
			Right: nil,
		},
	}
	expected := 2

	if reflect.DeepEqual(expected, minDepth(params.tree)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test4(t *testing.T) {
	params := parameters{
		tree: &TreeNode{Val: 1,
			Left: &TreeNode{Val: 2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 5}},
			Right: &TreeNode{Val: 3},
		},
	}
	expected := 2

	if reflect.DeepEqual(expected, minDepth(params.tree)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test5(t *testing.T) {
	params := parameters{
		tree: &TreeNode{Val: -9,
			Left: &TreeNode{Val: -3,
				Right: &TreeNode{Val: 4,
					Left: &TreeNode{Val: -6}}},
			Right: &TreeNode{Val: 2,
				Left: &TreeNode{Val: 4,
					Left: &TreeNode{Val: -5}},
				Right: &TreeNode{Val: 0}},
		},
	}
	expected := 3

	if reflect.DeepEqual(expected, minDepth(params.tree)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

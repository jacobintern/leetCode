package Q25

import (
	"reflect"
	"testing"
)

type parameters struct {
	tree      *TreeNode
	targetSum int
}

func Test1(t *testing.T) {
	params := parameters{
		tree: &TreeNode{Val: 5,
			Left: &TreeNode{Val: 4,
				Left: &TreeNode{Val: 11,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 2},
				},
			},
			Right: &TreeNode{Val: 8,
				Left: &TreeNode{Val: 13},
				Right: &TreeNode{Val: 4,
					Right: &TreeNode{Val: 1},
				},
			},
		},
		targetSum: 22,
	}
	expected := true

	if reflect.DeepEqual(expected, hasPathSum(params.tree, params.targetSum)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		tree: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 3},
		},
		targetSum: 5,
	}
	expected := false

	if reflect.DeepEqual(expected, hasPathSum(params.tree, params.targetSum)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test3(t *testing.T) {
	params := parameters{
		tree: &TreeNode{Val: 1,
			Left: &TreeNode{Val: 2},
		},
		targetSum: 0,
	}
	expected := false

	if reflect.DeepEqual(expected, hasPathSum(params.tree, params.targetSum)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

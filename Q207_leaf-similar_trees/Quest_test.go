package Q207

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
		root1: &TreeNode{Val: 3,
			Left: &TreeNode{Val: 5,
				Left: &TreeNode{Val: 6},
				Right: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 4}}},
			Right: &TreeNode{Val: 1,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 8}},
		},
		root2: &TreeNode{Val: 3,
			Left: &TreeNode{Val: 5,
				Left:  &TreeNode{Val: 6},
				Right: &TreeNode{Val: 7}},
			Right: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 4},
				Right: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 9},
					Right: &TreeNode{Val: 8}}},
		},
	}

	expected := true

	if testResult := leafSimilar(params.root1, params.root2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {
	params := parameters{
		root1: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 3},
		},
		root2: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 2},
		},
	}

	expected := false

	if testResult := leafSimilar(params.root1, params.root2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

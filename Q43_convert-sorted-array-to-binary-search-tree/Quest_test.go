package Q43

import (
	"reflect"
	"testing"
)

type parameters struct {
	arr []int
}

func Test1(t *testing.T) {

	params := parameters{arr: []int{-10, -3, 0, 5, 9}}

	expected := &TreeNode{Val: 0,
		Left: &TreeNode{Val: -10,
			Left:  nil,
			Right: &TreeNode{Val: -3},
		},
		Right: &TreeNode{Val: 5,
			Left:  nil,
			Right: &TreeNode{Val: 9},
		},
	}

	if reflect.DeepEqual(expected, sortedArrayToBST(params.arr)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {

	params := parameters{arr: []int{1, 3}}

	expected := &TreeNode{Val: 1,
		Left:  nil,
		Right: &TreeNode{Val: 3},
	}

	if reflect.DeepEqual(expected, sortedArrayToBST(params.arr)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

package Q144

import (
	"reflect"
	"testing"
)

type parameters struct {
	root *TreeNode
}

func Test1(t *testing.T) {
	params := parameters{root: &TreeNode{Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}}

	expected := 1

	if testResult := getMinimumDifference(params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {
	params := parameters{root: &TreeNode{Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 48,
			Left: &TreeNode{
				Val: 12,
			},
			Right: &TreeNode{
				Val: 49,
			},
		},
	}}

	expected := 1

	if testResult := getMinimumDifference(params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {
	params := parameters{root: &TreeNode{Val: 1,
		Left: nil,
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}}

	expected := 2

	if testResult := getMinimumDifference(params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {
	params := parameters{root: &TreeNode{Val: 236,
		Left: &TreeNode{
			Val: 104,
			Right: &TreeNode{
				Val: 227,
			},
		},
		Right: &TreeNode{
			Val: 701,
			Right: &TreeNode{
				Val: 911,
			},
		},
	}}

	expected := 9

	if testResult := getMinimumDifference(params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

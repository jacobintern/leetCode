package Q131

import (
	"reflect"
	"testing"
)

type parameters struct {
	root *TreeNode
}

func Test1(t *testing.T) {

	params := parameters{root: &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7}}}}

	expected := []float64{3, 14.5, 11}

	if testResult := averageOfLevels(params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{root: &TreeNode{
		Val: 3,
		Left: &TreeNode{Val: 9,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7}},
		Right: &TreeNode{Val: 20}}}

	expected := []float64{3, 14.5, 11}

	if testResult := averageOfLevels(params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

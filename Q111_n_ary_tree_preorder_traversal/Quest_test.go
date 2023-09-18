package Q111

import (
	"reflect"
	"testing"
)

type parameters struct {
	root Node
}

func Test1(t *testing.T) {

	params := parameters{root: Node{Val: 1,
		Children: []*Node{
			{Val: 3,
				Children: []*Node{
					{Val: 5, Children: []*Node{}},
					{Val: 6, Children: []*Node{}},
				}},
			{Val: 2, Children: []*Node{}},
			{Val: 4, Children: []*Node{}},
		}}}

	expected := []int{1, 3, 5, 6, 2, 4}

	if testResult := preorder(&params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{root: Node{Val: 1, Children: []*Node{
		{Val: 2, Children: []*Node{}},
		{Val: 3, Children: []*Node{
			{Val: 6, Children: []*Node{}},
			{Val: 7, Children: []*Node{
				{Val: 11, Children: []*Node{
					{Val: 14, Children: []*Node{}},
				}},
			}},
		}},
		{Val: 4, Children: []*Node{
			{Val: 8, Children: []*Node{
				{Val: 12, Children: []*Node{}},
			}},
		}},
		{Val: 5, Children: []*Node{
			{Val: 9, Children: []*Node{
				{Val: 13, Children: []*Node{}},
			}},
			{Val: 10, Children: []*Node{}},
		}},
	}}}

	expected := []int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10}

	if testResult := preorder(&params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

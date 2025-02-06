package Q112

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

	expected := []int{5, 6, 3, 2, 4, 1}

	if testResult := postorder(&params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
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

	expected := []int{2, 6, 14, 11, 7, 3, 12, 8, 4, 13, 9, 10, 5, 1}

	if testResult := postorder(&params.root); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

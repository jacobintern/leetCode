package Q106

import (
	"reflect"
	"testing"
)

type parameters struct {
	node Node
}

func Test1(t *testing.T) {

	params := parameters{node: Node{Val: 1, Children: []*Node{
		&Node{Val: 3, Children: []*Node{
			&Node{Val: 5, Children: nil},
			&Node{Val: 6, Children: nil}}},
		&Node{Val: 2, Children: nil},
		&Node{Val: 4, Children: nil},
	}}}

	expected := 3

	if testResult := maxDepth(&params.node); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{node: Node{Val: 1, Children: []*Node{
		&Node{Val: 2, Children: nil},
		&Node{Val: 3, Children: []*Node{
			&Node{Val: 6, Children: nil},
			&Node{Val: 7, Children: []*Node{
				&Node{Val: 11, Children: []*Node{
					&Node{Val: 14, Children: nil}}}}}}},
		&Node{Val: 4, Children: []*Node{
			&Node{Val: 8, Children: []*Node{
				&Node{Val: 12, Children: nil}}}}},
		&Node{Val: 5, Children: []*Node{
			&Node{Val: 9, Children: []*Node{
				&Node{Val: 13, Children: nil}}},
			&Node{Val: 10, Children: nil}},
		}}}}

	expected := 5

	if testResult := maxDepth(&params.node); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

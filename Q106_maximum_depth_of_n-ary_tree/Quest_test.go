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
		{Val: 3, Children: []*Node{
			{Val: 5, Children: nil},
			{Val: 6, Children: nil}}},
		{Val: 2, Children: nil},
		{Val: 4, Children: nil},
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
		{Val: 2, Children: nil},
		{Val: 3, Children: []*Node{
			{Val: 6, Children: nil},
			{Val: 7, Children: []*Node{
				{Val: 11, Children: []*Node{
					{Val: 14, Children: nil}}}}}}},
		{Val: 4, Children: []*Node{
			{Val: 8, Children: []*Node{
				{Val: 12, Children: nil}}}}},
		{Val: 5, Children: []*Node{
			{Val: 9, Children: []*Node{
				{Val: 13, Children: nil}}},
			{Val: 10, Children: nil}},
		}}}}

	expected := 5

	if testResult := maxDepth(&params.node); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

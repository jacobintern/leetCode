package Q65

import (
	"reflect"
	"testing"
)

type parameters struct {
	arr   []int
	left  int
	right int
}

func Test1(t *testing.T) {

	params := parameters{arr: []int{-2, 0, 3, -5, 2, -1}, left: 0, right: 2}
	obj := Constructor(params.arr)

	expected := 1

	if testResult := obj.SumRange(params.left, params.right); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{arr: []int{-2, 0, 3, -5, 2, -1}, left: 2, right: 5}
	obj := Constructor(params.arr)

	expected := -1

	if testResult := obj.SumRange(params.left, params.right); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{arr: []int{-2, 0, 3, -5, 2, -1}, left: 0, right: 5}
	obj := Constructor(params.arr)

	expected := -3

	if testResult := obj.SumRange(params.left, params.right); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

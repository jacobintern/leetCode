package Q163

import (
	"reflect"
	"testing"
)

type parameters struct {
	matrix [][]int
}

func Test1(t *testing.T) {

	params := parameters{matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}

	expected := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}

	if testResult := spiralOrder(params.matrix); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}}

	expected := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}

	if testResult := spiralOrder(params.matrix); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

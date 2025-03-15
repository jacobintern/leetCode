package Q188

import (
	"reflect"
	"testing"
)

type parameters struct {
	matrix [][]int
}

func Test1(t *testing.T) {

	params := parameters{matrix: [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}}

	expected := true

	if testResult := isToeplitzMatrix(params.matrix); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{matrix: [][]int{{1, 2}, {2, 2}}}

	expected := false

	if testResult := isToeplitzMatrix(params.matrix); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

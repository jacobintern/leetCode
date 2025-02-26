package Q161

import (
	"reflect"
	"testing"
)

type parameters struct {
	matrix [][]int
}

func Test1(t *testing.T) {

	params := parameters{matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}

	expected := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}

	rotate(params.matrix)

	if testResult := params.matrix; reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{matrix: [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}}

	expected := [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}

	rotate(params.matrix)

	if testResult := params.matrix; reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

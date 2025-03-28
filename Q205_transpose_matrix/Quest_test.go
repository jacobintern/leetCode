package Q205

import (
	"reflect"
	"testing"
)

type parameters struct {
	matrix [][]int
}

func Test1(t *testing.T) {

	params := parameters{matrix: [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}}

	expected := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}

	if testResult := transpose(params.matrix); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{matrix: [][]int{{1, 2, 3}, {4, 5, 6}}}

	expected := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}

	if testResult := transpose(params.matrix); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

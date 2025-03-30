package Q209

import (
	"reflect"
	"testing"
)

type parameters struct {
	grid [][]int
}

func Test1(t *testing.T) {

	params := parameters{grid: [][]int{{1, 2}, {3, 4}}}

	expected := 17

	if testResult := projectionArea(params.grid); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{grid: [][]int{{2}}}

	expected := 5

	if testResult := projectionArea(params.grid); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{grid: [][]int{{1, 0}, {0, 2}}}

	expected := 8

	if testResult := projectionArea(params.grid); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

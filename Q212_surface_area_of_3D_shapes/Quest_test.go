package Q212

import (
	"reflect"
	"testing"
)

type parameters struct {
	grid [][]int
}

func Test1(t *testing.T) {

	params := parameters{grid: [][]int{{1, 2}, {3, 4}}}

	expected := 34

	if testResult := surfaceArea(params.grid); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{grid: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}}

	expected := 32

	if testResult := surfaceArea(params.grid); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{grid: [][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}}

	expected := 46

	if testResult := surfaceArea(params.grid); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

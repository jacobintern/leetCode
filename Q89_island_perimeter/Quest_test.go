package Q89

import (
	"reflect"
	"testing"
)

type parameters struct {
	grid [][]int
}

func Test1(t *testing.T) {

	params := parameters{grid: [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}}

	expected := 16

	if testResult := islandPerimeter(params.grid); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{grid: [][]int{{1}}}

	expected := 4

	if testResult := islandPerimeter(params.grid); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{grid: [][]int{{1, 0}}}

	expected := 4

	if testResult := islandPerimeter(params.grid); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{grid: [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 1, 1}}}

	expected := 20

	if testResult := islandPerimeter(params.grid); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

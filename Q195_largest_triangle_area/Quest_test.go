package Q195

import (
	"reflect"
	"testing"
)

type parameters struct {
	points [][]int
}

func Test1(t *testing.T) {

	params := parameters{points: [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}}

	expected := 2.0

	if testResult := largestTriangleArea(params.points); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{points: [][]int{{1, 0}, {0, 0}, {0, 1}}}

	expected := 0.5

	if testResult := largestTriangleArea(params.points); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{points: [][]int{{4, 6}, {6, 5}, {3, 1}}}

	expected := 5.5

	if testResult := largestTriangleArea(params.points); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

package Q165

import (
	"reflect"
	"testing"
)

type parameters struct {
	intervals [][]int
}

func Test1(t *testing.T) {

	params := parameters{intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}}

	expected := [][]int{{1, 6}, {8, 10}, {15, 18}}

	if testResult := merge(params.intervals); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{intervals: [][]int{{1, 4}, {4, 5}}}

	expected := [][]int{{1, 5}}

	if testResult := merge(params.intervals); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

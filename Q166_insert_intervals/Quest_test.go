package Q166

import (
	"reflect"
	"testing"
)

type parameters struct {
	intervals   [][]int
	newInterval []int
}

func Test1(t *testing.T) {

	params := parameters{
		intervals:   [][]int{{1, 3}, {6, 9}},
		newInterval: []int{2, 5},
	}

	expected := [][]int{{1, 5}, {6, 9}}

	if testResult := insert(params.intervals, params.newInterval); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{
		intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
		newInterval: []int{4, 8},
	}

	expected := [][]int{{1, 2}, {3, 10}, {12, 16}}

	if testResult := insert(params.intervals, params.newInterval); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

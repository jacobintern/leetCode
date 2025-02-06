package Q116

import (
	"reflect"
	"testing"
)

type parameters struct {
	m, n int
	ops  [][]int
}

func Test1(t *testing.T) {

	params := parameters{m: 3, n: 3, ops: [][]int{{2, 2}, {3, 3}}}

	expected := 4

	if testResult := maxCount(params.m, params.n, params.ops); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{m: 3, n: 3, ops: [][]int{{2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}}}

	expected := 4

	if testResult := maxCount(params.m, params.n, params.ops); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{m: 3, n: 3, ops: [][]int{}}

	expected := 9

	if testResult := maxCount(params.m, params.n, params.ops); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

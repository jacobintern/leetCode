package Q167

import (
	"reflect"
	"testing"
)

type parameters struct {
	m, n int
}

func Test1(t *testing.T) {

	params := parameters{m: 7, n: 3}

	expected := 28

	if testResult := uniquePaths(params.m, params.n); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{m: 3, n: 2}

	expected := 3

	if testResult := uniquePaths(params.m, params.n); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

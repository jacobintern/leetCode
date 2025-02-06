package Q109

import (
	"reflect"
	"testing"
)

type parameters struct {
	num2s [][]int
	r     int
	c     int
}

func Test1(t *testing.T) {

	params := parameters{num2s: [][]int{{1, 2}, {3, 4}}, r: 1, c: 4}

	expected := [][]int{{1, 2, 3, 4}}

	if testResult := matrixReshape(params.num2s, params.r, params.c); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{num2s: [][]int{{1, 2}, {3, 4}}, r: 2, c: 4}

	expected := [][]int{{1, 2}, {3, 4}}

	if testResult := matrixReshape(params.num2s, params.r, params.c); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

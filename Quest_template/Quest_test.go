package Q201

import (
	"reflect"
	"testing"
)

type parameters struct {
	rec1, rec2 []int
}

func Test1(t *testing.T) {

	params := parameters{
		rec1: []int{0, 0, 2, 2},
		rec2: []int{1, 1, 3, 3},
	}

	expected := true

	if testResult := isRectangleOverlap(params.rec1, params.rec2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{
		rec1: []int{0, 0, 2, 2},
		rec2: []int{1, 0, 2, 1},
	}

	expected := false

	if testResult := isRectangleOverlap(params.rec1, params.rec2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{
		rec1: []int{0, 0, 2, 2},
		rec2: []int{2, 2, 3, 3},
	}

	expected := false

	if testResult := isRectangleOverlap(params.rec1, params.rec2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

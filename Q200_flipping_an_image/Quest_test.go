package Q200

import (
	"reflect"
	"testing"
)

type parameters struct {
	image [][]int
}

func Test1(t *testing.T) {

	params := parameters{image: [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}}

	expected := [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}

	if testResult := flipAndInvertImage(params.image); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{image: [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}}

	expected := [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}

	if testResult := flipAndInvertImage(params.image); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

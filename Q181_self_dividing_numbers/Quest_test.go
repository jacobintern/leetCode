package Q181

import (
	"reflect"
	"testing"
)

type parameters struct {
	left, right int
}

func Test1(t *testing.T) {

	params := parameters{left: 1, right: 22}

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}

	if testResult := selfDividingNumbers(params.left, params.right); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{left: 47, right: 85}

	expected := []int{48, 55, 66, 77}

	if testResult := selfDividingNumbers(params.left, params.right); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

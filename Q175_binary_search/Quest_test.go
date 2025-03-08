package Q175

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums   []int
	target int
}

func Test1(t *testing.T) {

	params := parameters{
		nums:   []int{-1, 0, 3, 5, 9, 12},
		target: 9,
	}

	expected := 4

	if testResult := search(params.nums, params.target); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{
		nums:   []int{-1, 0, 3, 5, 9, 12},
		target: 2,
	}

	expected := -1

	if testResult := search(params.nums, params.target); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

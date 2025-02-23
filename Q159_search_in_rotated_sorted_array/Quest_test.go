package Q159

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums   []int
	target int
}

func Test1(t *testing.T) {

	params := parameters{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0}

	expected := 4

	if testResult := search(params.nums, params.target); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3}

	expected := -1

	if testResult := search(params.nums, params.target); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{nums: []int{1}, target: 0}

	expected := -1

	if testResult := search(params.nums, params.target); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

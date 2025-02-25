package Q160

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums   []int
	target int
}

func Test1(t *testing.T) {

	params := parameters{nums: []int{1, 2, 3}, target: 4}

	expected := 7
	// 1 1 1 1
	// 1 1 2
	// 1 2 1
	// 1 3
	// 2 1 1
	// 2 2
	// 3 1

	if testResult := combinationSum4(params.nums, params.target); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{nums: []int{9}, target: 3}

	expected := 0

	if testResult := combinationSum4(params.nums, params.target); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

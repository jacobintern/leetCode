package Q190

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums []int
	k    int
}

func Test1(t *testing.T) {

	params := parameters{nums: []int{2, 3, 5, 9}, k: 2}

	expected := 2

	if testResult := minCapability(params.nums, params.k); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{nums: []int{2, 7, 9, 3, 1}, k: 2}

	expected := 2

	if testResult := minCapability(params.nums, params.k); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

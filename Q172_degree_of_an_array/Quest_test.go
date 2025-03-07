package Q172

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums []int
}

func Test1(t *testing.T) {

	params := parameters{nums: []int{1, 2, 2, 3, 1}}

	expected := 2

	if testResult := findShortestSubArray(params.nums); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{nums: []int{1, 2, 2, 3, 1, 4, 2}}

	expected := 6

	if testResult := findShortestSubArray(params.nums); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{nums: []int{2, 1}}

	expected := 1

	if testResult := findShortestSubArray(params.nums); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

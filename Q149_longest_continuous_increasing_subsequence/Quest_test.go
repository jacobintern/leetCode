package Q149

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums []int
}

func Test1(t *testing.T) {

	params := parameters{nums: []int{1, 3, 5, 4, 7}}

	expected := 3

	if testResult := findLengthOfLCIS(params.nums); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{nums: []int{2, 2, 2, 2, 2}}

	expected := 1

	if testResult := findLengthOfLCIS(params.nums); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{nums: []int{1, 3, 5, 4, 2, 3, 4, 5}}

	expected := 4

	if testResult := findLengthOfLCIS(params.nums); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

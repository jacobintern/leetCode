package Q93

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums []int
	num  int
}

func Test1(t *testing.T) {

	params := parameters{nums: []int{1, 4}, num: 2}

	expected := 4

	if testResult := findPoisonedDuration(params.nums, params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{nums: []int{1, 2}, num: 2}

	expected := 3

	if testResult := findPoisonedDuration(params.nums, params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{nums: []int{1, 6, 7, 9}, num: 3}

	expected := 9

	if testResult := findPoisonedDuration(params.nums, params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, num: 10}

	expected := 18

	if testResult := findPoisonedDuration(params.nums, params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

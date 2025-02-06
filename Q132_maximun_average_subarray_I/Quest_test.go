package Q132

import (
	"reflect"
	"testing"
)

type parameters struct {
	arr []int
	k   int
}

func Test1(t *testing.T) {

	params := parameters{arr: []int{1, 12, -5, -6, 50, 3}, k: 4}

	expected := 12.75

	if testResult := findMaxAverage(params.arr, params.k); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{arr: []int{5}, k: 1}

	expected := 5.00

	if testResult := findMaxAverage(params.arr, params.k); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{arr: []int{0, 4, 0, 3, 2}, k: 1}

	expected := 4.00

	if testResult := findMaxAverage(params.arr, params.k); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

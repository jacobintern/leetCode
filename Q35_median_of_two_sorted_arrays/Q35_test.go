package Q35

import (
	"reflect"
	"testing"
)

type parameters struct {
	arr1 []int
	arr2 []int
}

func Test1(t *testing.T) {

	params := parameters{arr1: []int{1, 3}, arr2: []int{2}}

	expected := 2.00000

	if testResult := findMedianSortedArrays(params.arr1, params.arr2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{arr1: []int{1, 2}, arr2: []int{3, 4}}

	expected := 2.50000

	if testResult := findMedianSortedArrays(params.arr1, params.arr2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{arr1: []int{0, 0}, arr2: []int{0, 0}}

	expected := 0.00000

	if testResult := findMedianSortedArrays(params.arr1, params.arr2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{arr1: []int{}, arr2: []int{1}}

	expected := 1.00000

	if testResult := findMedianSortedArrays(params.arr1, params.arr2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test5(t *testing.T) {

	params := parameters{arr1: []int{2}, arr2: []int{}}

	expected := 2.00000

	if testResult := findMedianSortedArrays(params.arr1, params.arr2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

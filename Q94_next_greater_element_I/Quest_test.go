package Q94

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums1 []int
	nums2 []int
}

func Test1(t *testing.T) {

	params := parameters{nums1: []int{4, 1, 2}, nums2: []int{1, 3, 4, 2}}

	expected := []int{-1, 3, -1}

	if testResult := nextGreaterElement(params.nums1, params.nums2); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{nums1: []int{2, 4}, nums2: []int{1, 2, 3, 4}}

	expected := []int{3, -1}

	if testResult := nextGreaterElement(params.nums1, params.nums2); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{nums1: []int{1, 3, 5, 2, 4}, nums2: []int{6, 5, 4, 3, 2, 1, 7}}

	expected := []int{7, 7, 7, 7, 7}

	if testResult := nextGreaterElement(params.nums1, params.nums2); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

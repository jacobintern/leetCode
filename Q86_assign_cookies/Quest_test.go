package Q86

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums1 []int
	nums2 []int
}

func Test1(t *testing.T) {

	params := parameters{nums1: []int{1, 2, 3}, nums2: []int{1, 1}}

	expected := 1

	if testResult := findContentChildren(params.nums1, params.nums2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{nums1: []int{1, 2}, nums2: []int{1, 2, 3}}

	expected := 2

	if testResult := findContentChildren(params.nums1, params.nums2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

package Q127

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums1 []int
	nums2 []int
}

func Test1(t *testing.T) {

	params := parameters{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}}

	expected := []int{2, 2}

	if testResult := intersect(params.nums1, params.nums2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}}

	expected := []int{4, 9}

	if testResult := intersect(params.nums1, params.nums2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

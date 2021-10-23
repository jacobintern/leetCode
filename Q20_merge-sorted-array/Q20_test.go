package Q20

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
}

func Test1(t *testing.T) {
	params := parameters{
		nums1: []int{1, 2, 3, 0, 0, 0},
		m:     3,
		nums2: []int{2, 5, 6},
		n:     3,
	}
	expected := []int{1, 2, 2, 3, 5, 6}

	if reflect.DeepEqual(expected, merge(params.nums1, params.m, params.nums2, params.n)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		nums1: []int{1},
		m:     1,
		nums2: []int{},
		n:     0,
	}
	expected := []int{1}

	if reflect.DeepEqual(expected, merge(params.nums1, params.m, params.nums2, params.n)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test3(t *testing.T) {
	params := parameters{
		nums1: []int{0},
		m:     0,
		nums2: []int{1},
		n:     1,
	}
	expected := []int{1}

	if reflect.DeepEqual(expected, merge(params.nums1, params.m, params.nums2, params.n)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

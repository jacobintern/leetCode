package Q11

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums []int
	num  int
}

func Test1(t *testing.T) {
	params := parameters{
		nums: []int{1, 3, 5, 6},
		num:  5,
	}
	expected := 2

	if reflect.DeepEqual(expected, searchInsert(params.nums, params.num)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		nums: []int{1, 3, 5, 6},
		num:  2,
	}
	expected := 1

	if reflect.DeepEqual(expected, searchInsert(params.nums, params.num)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test3(t *testing.T) {
	params := parameters{
		nums: []int{1, 3, 5, 6},
		num:  7,
	}
	expected := 4

	if reflect.DeepEqual(expected, searchInsert(params.nums, params.num)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test4(t *testing.T) {
	params := parameters{
		nums: []int{1, 3, 5, 6},
		num:  0,
	}
	expected := 0

	if reflect.DeepEqual(expected, searchInsert(params.nums, params.num)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test5(t *testing.T) {
	params := parameters{
		nums: []int{1},
		num:  0,
	}
	expected := 0

	if reflect.DeepEqual(expected, searchInsert(params.nums, params.num)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test6(t *testing.T) {
	params := parameters{
		nums: []int{2, 5},
		num:  1,
	}
	expected := 0

	if reflect.DeepEqual(expected, searchInsert(params.nums, params.num)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

package Q13

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums []int
}

func Test1(t *testing.T) {
	params := parameters{
		nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
	}
	expected := 6

	if reflect.DeepEqual(expected, maxSubArray(params.nums)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		nums: []int{1},
	}
	expected := 1

	if reflect.DeepEqual(expected, maxSubArray(params.nums)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test3(t *testing.T) {
	params := parameters{
		nums: []int{5, 4, -1, 7, 8},
	}
	expected := 23

	if reflect.DeepEqual(expected, maxSubArray(params.nums)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

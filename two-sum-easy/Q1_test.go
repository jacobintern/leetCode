package Q1

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums   []int
	target int
}

func Test1(t *testing.T) {
	params := parameters{
		nums:   []int{2, 5, 7, 11, 15},
		target: 9,
	}
	var expected []int
	expected = append(expected, 0, 2)

	if reflect.DeepEqual(expected, twoSum(params.nums, params.target)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		nums:   []int{3, 2, 4},
		target: 6,
	}
	var expected []int
	expected = append(expected, 1, 2)

	if reflect.DeepEqual(expected, twoSum(params.nums, params.target)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

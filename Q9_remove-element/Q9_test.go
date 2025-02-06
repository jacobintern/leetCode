package Q9

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
		nums: []int{3, 2, 2, 3},
		num:  3,
	}
	expected := 2

	if reflect.DeepEqual(expected, removeElement(params.nums, params.num)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
		num:  2,
	}
	expected := 5

	if reflect.DeepEqual(expected, removeElement(params.nums, params.num)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

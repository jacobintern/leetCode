package Q15

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums []int
}

func Test1(t *testing.T) {
	params := parameters{
		nums: []int{1, 2, 3},
	}
	expected := []int{1, 2, 4}

	if reflect.DeepEqual(expected, plusOne(params.nums)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		nums: []int{4, 3, 2, 1},
	}
	expected := []int{4, 3, 2, 2}

	if reflect.DeepEqual(expected, plusOne(params.nums)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test3(t *testing.T) {
	params := parameters{
		nums: []int{0},
	}
	expected := []int{1}

	if reflect.DeepEqual(expected, plusOne(params.nums)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test4(t *testing.T) {
	params := parameters{
		nums: []int{9, 9},
	}
	expected := []int{1, 0, 0}

	if reflect.DeepEqual(expected, plusOne(params.nums)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

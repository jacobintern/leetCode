package Q8

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums []int
}

func Test1(t *testing.T) {
	params := parameters{
		nums: []int{1, 1, 2},
	}
	expected := 2
	// nums = [1,2]

	if reflect.DeepEqual(expected, removeDuplicates2(params.nums)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
	}
	expected := 5
	// nums = [0,1,2,3,4]

	if reflect.DeepEqual(expected, removeDuplicates2(params.nums)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

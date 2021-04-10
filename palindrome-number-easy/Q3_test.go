package Q3

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums int
}

func Test1(t *testing.T) {
	params := parameters{
		nums: 212,
	}
	expected := true

	if reflect.DeepEqual(expected, IsPalindrome(params.nums)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		nums: -10,
	}
	expected := false

	if reflect.DeepEqual(expected, IsPalindrome(params.nums)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

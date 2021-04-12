package Q3

import (
	"reflect"
	"testing"
)

type parameters struct {
	num int
}

func Test1(t *testing.T) {
	params := parameters{
		num: 212,
	}
	expected := true

	if reflect.DeepEqual(expected, IsPalindrome(params.num)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		num: -10,
	}
	expected := false

	if reflect.DeepEqual(expected, IsPalindrome(params.num)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

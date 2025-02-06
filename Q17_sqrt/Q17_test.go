package Q17

import (
	"reflect"
	"testing"
)

type parameters struct {
	num int
	b   string
}

func Test1(t *testing.T) {
	params := parameters{
		num: 4,
	}
	expected := 2

	if reflect.DeepEqual(expected, mySqrt(params.num)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		num: 8,
	}
	expected := 2

	if reflect.DeepEqual(expected, mySqrt(params.num)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

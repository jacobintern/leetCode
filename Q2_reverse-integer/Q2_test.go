package Q2

import (
	"reflect"
	"testing"
)

type parameters struct {
	num int
}

func Test1(t *testing.T) {
	params := parameters{
		num: -213,
	}
	expected := -312

	if reflect.DeepEqual(expected, reverse(params.num)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

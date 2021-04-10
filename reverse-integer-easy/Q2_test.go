package Q2

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums int
}

func Test1(t *testing.T) {
	params := parameters{
		nums: -213,
	}
	expected := -312

	if reflect.DeepEqual(expected, reverse(params.nums)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

package Q18

import (
	"reflect"
	"testing"
)

type parameters struct {
	num int
}

func Test1(t *testing.T) {
	params := parameters{
		num: 2,
	}
	expected := 2

	if reflect.DeepEqual(expected, climbStairs(params.num)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		num: 3,
	}
	expected := 3

	if reflect.DeepEqual(expected, climbStairs(params.num)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

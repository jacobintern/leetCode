package Q12

import (
	"reflect"
	"testing"
)

type parameters struct {
	x float64
	n int
}

func Test1(t *testing.T) {
	params := parameters{
		x: 2.00000,
		n: 10,
	}
	expected := 1024.00000

	if reflect.DeepEqual(expected, myPow(params.x, params.n)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		x: 2.10000,
		n: 3,
	}
	expected := 9.26100

	if reflect.DeepEqual(expected, myPow(params.x, params.n)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test3(t *testing.T) {
	params := parameters{
		x: 2.00000,
		n: -2,
	}
	expected := 0.25000

	if reflect.DeepEqual(expected, myPow(params.x, params.n)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test4(t *testing.T) {
	params := parameters{
		x: 1.00000,
		n: -2147483648,
	}
	expected := 1

	if reflect.DeepEqual(expected, myPow(params.x, params.n)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

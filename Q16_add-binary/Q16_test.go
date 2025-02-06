package Q16

import (
	"reflect"
	"testing"
)

type parameters struct {
	a string
	b string
}

func Test1(t *testing.T) {
	params := parameters{
		a: "11",
		b: "1",
	}
	expected := "100"

	if reflect.DeepEqual(expected, addBinary(params.a, params.b)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		a: "1010",
		b: "1011",
	}
	expected := "10101"

	if reflect.DeepEqual(expected, addBinary(params.a, params.b)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test3(t *testing.T) {
	params := parameters{
		a: "11",
		b: "11",
	}
	expected := "110"

	if reflect.DeepEqual(expected, addBinary(params.a, params.b)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test4(t *testing.T) {
	params := parameters{
		a: "1",
		b: "111",
	}
	expected := "1000"

	if reflect.DeepEqual(expected, addBinary(params.a, params.b)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

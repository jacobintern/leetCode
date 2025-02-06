package Q10

import (
	"reflect"
	"testing"
)

type parameters struct {
	str    string
	needle string
}

func Test1(t *testing.T) {
	params := parameters{
		str:    "hello",
		needle: "ll",
	}
	expected := 2

	if reflect.DeepEqual(expected, strStr(params.str, params.needle)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		str:    "aaaaa",
		needle: "bba",
	}
	expected := -1

	if reflect.DeepEqual(expected, strStr(params.str, params.needle)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test3(t *testing.T) {
	params := parameters{
		str:    "",
		needle: "",
	}
	expected := 0

	if reflect.DeepEqual(expected, strStr(params.str, params.needle)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test4(t *testing.T) {
	params := parameters{
		str:    "a",
		needle: "a",
	}
	expected := 0

	if reflect.DeepEqual(expected, strStr(params.str, params.needle)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

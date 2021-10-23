package Q4

import (
	"reflect"
	"testing"
)

type parameters struct {
	str string
}

func Test1(t *testing.T) {
	params := parameters{
		str: "MCMXCIV",
	}
	expected := 1994

	if reflect.DeepEqual(expected, romanToInt(params.str)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		str: "III",
	}
	expected := 3

	if reflect.DeepEqual(expected, romanToInt(params.str)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test3(t *testing.T) {
	params := parameters{
		str: "IV",
	}
	expected := 4

	if reflect.DeepEqual(expected, romanToInt(params.str)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test4(t *testing.T) {
	params := parameters{
		str: "IX",
	}
	expected := 9

	if reflect.DeepEqual(expected, romanToInt(params.str)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

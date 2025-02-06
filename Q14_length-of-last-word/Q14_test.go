package Q14

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
		str: "hello wrold",
	}
	expected := 5

	if reflect.DeepEqual(expected, lengthOfLastWord(params.str)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		str: " ",
	}
	expected := 0

	if reflect.DeepEqual(expected, lengthOfLastWord(params.str)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test3(t *testing.T) {
	params := parameters{
		str: "a ",
	}
	expected := 1

	if reflect.DeepEqual(expected, lengthOfLastWord(params.str)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

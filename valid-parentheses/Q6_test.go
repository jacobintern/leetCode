package Q6

import (
	"reflect"
	"testing"
)

type parameters struct {
	str string
}

func Test1(t *testing.T) {
	params := parameters{
		str: "()",
	}
	expected := true

	if reflect.DeepEqual(expected, isValid(params.str)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		str: "()[]{}",
	}
	expected := true

	if reflect.DeepEqual(expected, isValid(params.str)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test3(t *testing.T) {
	params := parameters{
		str: "(]",
	}
	expected := false

	if reflect.DeepEqual(expected, isValid(params.str)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test4(t *testing.T) {
	params := parameters{
		str: "([)]",
	}
	expected := false

	if reflect.DeepEqual(expected, isValid(params.str)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test5(t *testing.T) {
	params := parameters{
		str: "{[]}",
	}
	expected := true

	if reflect.DeepEqual(expected, isValid(params.str)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test6(t *testing.T) {
	params := parameters{
		str: "[ [ [ ] ] ] ",
	}
	expected := true

	if reflect.DeepEqual(expected, isValid(params.str)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

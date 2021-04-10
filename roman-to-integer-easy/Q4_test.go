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
		str: "III",
	}
	expected := 3

	if reflect.DeepEqual(expected, romanToInt(params.str)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

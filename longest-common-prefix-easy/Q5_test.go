package Q5

import (
	"reflect"
	"testing"
)

type parameters struct {
	strArr []string
}

func Test1(t *testing.T) {
	params := parameters{
		strArr: []string{"flower", "flow", "flight"},
	}
	expected := "fl"

	if reflect.DeepEqual(expected, longestCommonPrefix(params.strArr)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		strArr: []string{"dog", "racecar", "car"},
	}
	expected := "fl"

	if reflect.DeepEqual(expected, longestCommonPrefix(params.strArr)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

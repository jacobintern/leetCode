package Q192

import (
	"reflect"
	"testing"
)

type parameters struct {
	str  string
	goal string
}

func Test1(t *testing.T) {

	params := parameters{
		str:  "abcde",
		goal: "cdeab",
	}

	expected := 0

	if testResult := rotateString(params.str, params.goal); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{
		str:  "abcde",
		goal: "abced",
	}

	expected := 0

	if testResult := rotateString(params.str, params.goal); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

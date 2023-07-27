package Q77

import (
	"reflect"
	"testing"
)

type parameters struct {
	str1 string
	str2 string
}

func Test1(t *testing.T) {

	params := parameters{str1: "acb", str2: "ahbgdc"}

	expected := false

	if testResult := isSubsequence(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{str1: "axc", str2: "ahbgdc"}

	expected := false

	if testResult := isSubsequence(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{str1: "b", str2: "abc"}

	expected := true

	if testResult := isSubsequence(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

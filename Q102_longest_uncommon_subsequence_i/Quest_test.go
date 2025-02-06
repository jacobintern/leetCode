package Q102

import (
	"reflect"
	"testing"
)

type parameters struct {
	str1 string
	str2 string
}

func Test1(t *testing.T) {

	params := parameters{str1: "aba", str2: "cdc"}

	expected := 3

	if testResult := findLUSlength(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{str1: "aaa", str2: "bbb"}

	expected := 3

	if testResult := findLUSlength(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{str1: "aaa", str2: "aaa"}

	expected := -1

	if testResult := findLUSlength(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

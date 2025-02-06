package Q103

import (
	"reflect"
	"testing"
)

type parameters struct {
	str string
	k   int
}

func Test1(t *testing.T) {

	params := parameters{str: "abcdefg", k: 2}

	expected := "bacdfeg"

	if testResult := reverseStr(params.str, params.k); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{str: "abcd", k: 2}

	expected := "bacd"

	if testResult := reverseStr(params.str, params.k); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{str: "abcdefg", k: 3}

	expected := "cbadefg"

	if testResult := reverseStr(params.str, params.k); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{str: "abcd", k: 4}

	expected := "dcba"

	if testResult := reverseStr(params.str, params.k); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test5(t *testing.T) {

	params := parameters{str: "abcdefg", k: 8}

	expected := "gfedcba"

	if testResult := reverseStr(params.str, params.k); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

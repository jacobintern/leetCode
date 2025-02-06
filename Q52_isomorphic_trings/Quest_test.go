package Q52

import (
	"reflect"
	"testing"
)

type parameters struct {
	str1 string
	str2 string
}

func Test1(t *testing.T) {

	params := parameters{str1: "add", str2: "egg"}

	expected := true

	if testResult := isIsomorphic(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{str1: "foo", str2: "bar"}

	expected := false

	if testResult := isIsomorphic(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{str1: "paper", str2: "title"}

	expected := true

	if testResult := isIsomorphic(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{str1: "bbbaaaba", str2: "aaabbbba"}

	expected := false

	if testResult := isIsomorphic(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test5(t *testing.T) {

	params := parameters{str1: "leel", str2: "lall"}

	expected := false

	if testResult := isIsomorphic(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test6(t *testing.T) {

	params := parameters{str1: "papap", str2: "titii"}

	expected := false

	if testResult := isIsomorphic(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test7(t *testing.T) {

	params := parameters{str1: "aaaa", str2: "aaaa"}

	expected := true

	if testResult := isIsomorphic(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test8(t *testing.T) {

	params := parameters{str1: "abcdefghijklmnopqrstuvwxyzva", str2: "abcdefghijklmnopqrstuvwxyzck"}

	expected := true

	if testResult := isIsomorphic(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

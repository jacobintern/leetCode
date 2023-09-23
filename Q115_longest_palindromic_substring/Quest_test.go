package Q115

import (
	"reflect"
	"testing"
)

type parameters struct {
	str string
}

func Test1(t *testing.T) {

	params := parameters{str: "babad"}

	expected := "bab"

	if testResult := longestPalindrome(params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{str: "cbbd"}

	expected := "bb"

	if testResult := longestPalindrome(params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{str: "ccc"}

	expected := "ccc"

	if testResult := longestPalindrome(params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{str: "xaabacxcabaaxcabaax"}

	expected := "xaabacxcabaax"

	if testResult := longestPalindrome(params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test5(t *testing.T) {

	params := parameters{str: "bananas"}

	expected := "anana"

	if testResult := longestPalindrome(params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

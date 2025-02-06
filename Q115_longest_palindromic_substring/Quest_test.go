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
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{str: "cbbd"}

	expected := "bb"

	if testResult := longestPalindrome(params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{str: "ccc"}

	expected := "ccc"

	if testResult := longestPalindrome(params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{str: "xaabacxcabaaxcabaax"}

	expected := "xaabacxcabaax"

	if testResult := longestPalindrome(params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test5(t *testing.T) {

	params := parameters{str: "bananas"}

	expected := "anana"

	if testResult := longestPalindrome(params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test6(t *testing.T) {

	params := parameters{str: "dbbd"}

	expected := "dbbd"

	if testResult := longestPalindrome(params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test7(t *testing.T) {

	params := parameters{str: "eabcb"}

	expected := "bcb"

	if testResult := longestPalindrome(params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test8(t *testing.T) {

	params := parameters{str: "aaabaaaa"}

	expected := "aaabaaa"

	if testResult := longestPalindrome(params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test9(t *testing.T) {

	params := parameters{str: "caaaaa"}

	expected := "aaaaa"

	if testResult := longestPalindrome(params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test10(t *testing.T) {

	params := parameters{str: "ababababa"}

	expected := "ababababa"

	if testResult := longestPalindrome(params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

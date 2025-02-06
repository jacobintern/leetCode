package Q63

import (
	"reflect"
	"testing"
)

type parameters struct {
	pattern string
	str     string
}

func Test1(t *testing.T) {

	params := parameters{pattern: "abba", str: "dog cat cat dog"}

	expected := true

	if testResult := wordPattern(params.pattern, params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{pattern: "abba", str: "dog cat cat fish"}

	expected := false

	if testResult := wordPattern(params.pattern, params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{pattern: "aaaa", str: "dog cat cat dog"}

	expected := false

	if testResult := wordPattern(params.pattern, params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{pattern: "abba", str: "dog dog dog dog"}

	expected := false

	if testResult := wordPattern(params.pattern, params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

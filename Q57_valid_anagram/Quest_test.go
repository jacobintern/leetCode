package Q57

import (
	"reflect"
	"testing"
)

type parameters struct {
	str1 string
	str2 string
}

func Test1(t *testing.T) {

	params := parameters{str1: "anagram", str2: "nagaram"}

	expected := true

	if testResult := isAnagram(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{str1: "rat", str2: "car"}

	expected := false

	if testResult := isAnagram(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

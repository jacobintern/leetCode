package Q203

import (
	"reflect"
	"testing"
)

type parameters struct {
	str, str2 string
}

func Test1(t *testing.T) {

	params := parameters{str: "ab", str2: "ba"}

	expected := true

	if testResult := buddyStrings(params.str, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{str: "ab", str2: "ab"}

	expected := false

	if testResult := buddyStrings(params.str, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}
func Test3(t *testing.T) {

	params := parameters{str: "aa", str2: "aa"}

	expected := true

	if testResult := buddyStrings(params.str, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

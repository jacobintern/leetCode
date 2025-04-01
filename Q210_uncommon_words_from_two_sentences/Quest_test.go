package Q210

import (
	"reflect"
	"testing"
)

type parameters struct {
	str, str2 string
}

func Test1(t *testing.T) {

	params := parameters{str: "this apple is sweet", str2: "this apple is sour"}

	expected := []string{"sweet", "sour"}

	if testResult := uncommonFromSentences(params.str, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{str: "apple apple", str2: "banana"}

	expected := []string{"banana"}

	if testResult := uncommonFromSentences(params.str, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

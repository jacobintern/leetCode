package Q196

import (
	"reflect"
	"testing"
)

type parameters struct {
	str    string
	banned []string
}

func Test1(t *testing.T) {

	params := parameters{
		str:    "Bob hit a ball, the hit BALL flew far after it was hit.",
		banned: []string{"hit"},
	}

	expected := "ball"

	if testResult := mostCommonWord(params.str, params.banned); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{
		str:    "a.",
		banned: []string{""},
	}

	expected := "a"

	if testResult := mostCommonWord(params.str, params.banned); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}


func Test3(t *testing.T) {

	params := parameters{
		str:    "Bob",
		banned: []string{""},
	}

	expected := "bob"

	if testResult := mostCommonWord(params.str, params.banned); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

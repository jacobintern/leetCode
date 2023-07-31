package Q83

import (
	"reflect"
	"testing"
)

type parameters struct {
	str1 string
	str2 string
}

func Test1(t *testing.T) {

	params := parameters{str1: "11", str2: "123"}

	expected := "134"

	if testResult := addStrings(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{str1: "456", str2: "77"}

	expected := "533"

	if testResult := addStrings(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{str1: "0", str2: "0"}

	expected := "0"

	if testResult := addStrings(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{str1: "9", str2: "9"}

	expected := "18"

	if testResult := addStrings(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test5(t *testing.T) {

	params := parameters{str1: "123456789", str2: "987654321"}

	expected := "1111111110"

	if testResult := addStrings(params.str1, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

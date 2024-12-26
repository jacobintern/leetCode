package Q143

import (
	"reflect"
	"testing"
)

type parameters struct {
	str string
	k   int
}

func Test1(t *testing.T) {

	params := parameters{str: "5F3Z-2e-9-wsoflgkoh", k: 4}

	expected := "5F3Z-2E9W-SOFL-GKOH"

	if testResult := licenseKeyFormatting(params.str, params.k); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{str: "2-5g-3-J", k: 2}

	expected := "2-5G-3J"

	if testResult := licenseKeyFormatting(params.str, params.k); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{str: "a-a-a-a-", k: 1}

	expected := "A-A-A-A"

	if testResult := licenseKeyFormatting(params.str, params.k); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{str: "a", k: 1}

	expected := "A"

	if testResult := licenseKeyFormatting(params.str, params.k); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test5(t *testing.T) {

	params := parameters{str: "aaaa", k: 2}

	expected := "AA-AA"

	if testResult := licenseKeyFormatting(params.str, params.k); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

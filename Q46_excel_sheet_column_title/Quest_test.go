package Q47

import (
	"reflect"
	"testing"
)

type parameters struct {
	target int
}

func Test1(t *testing.T) {

	params := parameters{target: 1}

	expected := "A"

	if testResult := convertToTitle(params.target); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{target: 28}

	expected := "AB"

	if testResult := convertToTitle(params.target); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{target: 701}

	expected := "ZY"

	if testResult := convertToTitle(params.target); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{target: 2147483647}

	expected := "FXSHRXW"

	if testResult := convertToTitle(params.target); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test5(t *testing.T) {

	params := parameters{target: 676}

	expected := "YZ"

	if testResult := convertToTitle(params.target); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

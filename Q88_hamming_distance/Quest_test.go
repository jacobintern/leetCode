package Q88

import (
	"reflect"
	"testing"
)

type parameters struct {
	num1 int
	num2 int
}

func Test1(t *testing.T) {

	params := parameters{num1: 1, num2: 4}

	expected := 2

	if testResult := hammingDistance(params.num1, params.num2); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{num1: 3, num2: 1}

	expected := 1

	if testResult := hammingDistance(params.num1, params.num2); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{num1: 3, num2: 3}

	expected := 0

	if testResult := hammingDistance(params.num1, params.num2); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

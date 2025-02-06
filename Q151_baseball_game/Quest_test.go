package Q151

import (
	"reflect"
	"testing"
)

type parameters struct {
	op []string
}

func Test1(t *testing.T) {

	params := parameters{op: []string{"5", "2", "C", "D", "+"}}

	expected := 30

	if testResult := calPoints(params.op); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{op: []string{"5", "-2", "4", "C", "D", "9", "+", "+"}}

	expected := 27

	if testResult := calPoints(params.op); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{op: []string{"1", "C"}}

	expected := 0

	if testResult := calPoints(params.op); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

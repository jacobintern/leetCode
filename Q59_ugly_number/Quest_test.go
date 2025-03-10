package Q59

import (
	"reflect"
	"testing"
)

type parameters struct {
	num int
}

func Test1(t *testing.T) {

	params := parameters{num: 6}

	expected := true

	if testResult := isUgly(params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{num: 1}

	expected := true

	if testResult := isUgly(params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{num: 14}

	expected := false

	if testResult := isUgly(params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

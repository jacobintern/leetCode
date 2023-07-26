package Q72

import (
	"reflect"
	"testing"
)

type parameters struct {
	num int
}

func Test1(t *testing.T) {

	params := parameters{num: 16}

	expected := true

	if testResult := isPerfectSquare(params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{num: 14}

	expected := false

	if testResult := isPerfectSquare(params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

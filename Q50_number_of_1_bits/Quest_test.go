package Q50

import (
	"reflect"
	"testing"
)

type parameters struct {
	num uint32
}

func Test1(t *testing.T) {

	params := parameters{num: 11}

	expected := 3

	if testResult := hammingWeight(params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{num: 128}

	expected := 1

	if testResult := hammingWeight(params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{num: 4294967293}

	expected := 31

	if testResult := hammingWeight(params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

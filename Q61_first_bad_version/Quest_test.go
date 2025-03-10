package Q61

import (
	"reflect"
	"testing"
)

type parameters struct {
	num int
}

func Test1(t *testing.T) {

	params := parameters{num: 5}

	expected := 4

	if testResult := firstBadVersion(params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{num: 1}

	expected := 1

	if testResult := firstBadVersion(params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

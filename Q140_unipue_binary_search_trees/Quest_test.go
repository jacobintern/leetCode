package Q140

import (
	"reflect"
	"testing"
)

type parameters struct {
	num int
}

func Test1(t *testing.T) {

	params := parameters{num: 3}

	expected := 5

	if testResult := numTrees(params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{num: 1}

	expected := 1

	if testResult := numTrees(params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

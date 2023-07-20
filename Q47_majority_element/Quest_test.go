package Q47

import (
	"reflect"
	"testing"
)

type parameters struct {
	arr []int
}

func Test1(t *testing.T) {

	params := parameters{arr: []int{3, 2, 3}}

	expected := 3

	if testResult := majorityElement(params.arr); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{arr: []int{2, 2, 1, 1, 1, 2, 2}}

	expected := 3

	if testResult := majorityElement(params.arr); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

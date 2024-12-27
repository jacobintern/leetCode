package Q146

import (
	"reflect"
	"testing"
)

type parameters struct {
	arr []int
}

func Test1(t *testing.T) {

	params := parameters{arr: []int{1, 2, 3}}

	expected := 6

	if testResult := maximumProduct(params.arr); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{arr: []int{1, 2, 3, 4}}

	expected := 24

	if testResult := maximumProduct(params.arr); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{arr: []int{-1, -2, -3}}

	expected := -6

	if testResult := maximumProduct(params.arr); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{arr: []int{29, 38, 32, 15, 83, -17, 51, -31, -56, -62, 79, -74, -48, -90, 45, -27, -51, 13, -24, -70}}

	expected := 552780

	if testResult := maximumProduct(params.arr); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

package Q37

import (
	"reflect"
	"testing"
)

type parameters struct {
	arr []int
}

func Test1(t *testing.T) {

	params := parameters{arr: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}}

	expected := 49

	if testResult := maxArea(params.arr); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{arr: []int{1, 1}}

	expected := 1

	if testResult := maxArea(params.arr); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{arr: []int{4, 3, 2, 1, 4}}

	expected := 16

	if testResult := maxArea(params.arr); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{arr: []int{1, 2, 1}}

	expected := 2

	if testResult := maxArea(params.arr); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

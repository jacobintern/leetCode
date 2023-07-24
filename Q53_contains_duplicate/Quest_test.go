package Q53

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums []int
}

func Test1(t *testing.T) {

	params := parameters{nums: []int{1, 2, 3, 1}}

	expected := true

	if testResult := containsDuplicate(params.nums); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{nums: []int{1, 2, 3, 4}}

	expected := false

	if testResult := containsDuplicate(params.nums); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}}

	expected := true

	if testResult := containsDuplicate(params.nums); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

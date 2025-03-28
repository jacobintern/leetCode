package Q204

import (
	"reflect"
	"testing"
)

type parameters struct {
	bills []int
}

func Test1(t *testing.T) {

	params := parameters{bills: []int{5, 5, 5, 10, 20}}

	expected := true

	if testResult := lemonadeChange(params.bills); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{bills: []int{5, 5, 10, 10, 20}}

	expected := false

	if testResult := lemonadeChange(params.bills); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

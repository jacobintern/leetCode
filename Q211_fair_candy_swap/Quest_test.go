package Q211

import (
	"reflect"
	"testing"
)

type parameters struct {
	alice, bob []int
}

func Test1(t *testing.T) {

	params := parameters{
		alice: []int{1, 1},
		bob:   []int{2, 2},
	}

	expected := []int{1, 2}

	if testResult := fairCandySwap(params.alice, params.bob); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{
		alice: []int{1, 2},
		bob:   []int{2, 3},
	}

	expected := []int{1, 2}

	if testResult := fairCandySwap(params.alice, params.bob); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{
		alice: []int{2},
		bob:   []int{1, 3},
	}

	expected := []int{2, 3}

	if testResult := fairCandySwap(params.alice, params.bob); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

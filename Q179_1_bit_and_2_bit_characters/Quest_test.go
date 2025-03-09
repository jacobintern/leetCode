package Q179

import (
	"reflect"
	"testing"
)

type parameters struct {
	bits []int
}

func Test1(t *testing.T) {

	params := parameters{bits: []int{1, 0, 0}}

	expected := true

	if testResult := isOneBitCharacter(params.bits); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{bits: []int{1, 1, 1, 0}}

	expected := false

	if testResult := isOneBitCharacter(params.bits); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

package Q183

import (
	"reflect"
	"testing"
)

type parameters struct {
	letters []byte
	target  byte
}

func Test1(t *testing.T) {

	params := parameters{letters: []byte{'c', 'f', 'j'}, target: 'a'}

	expected := byte('c')

	if testResult := nextGreatestLetter(params.letters, params.target); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{letters: []byte{'c', 'f', 'j'}, target: 'c'}

	expected := byte('f')

	if testResult := nextGreatestLetter(params.letters, params.target); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{letters: []byte{'x', 'x', 'y', 'y'}, target: 'z'}

	expected := byte('x')

	if testResult := nextGreatestLetter(params.letters, params.target); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

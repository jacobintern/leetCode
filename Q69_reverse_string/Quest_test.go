package Q69

import (
	"reflect"
	"testing"
)

type parameters struct {
	s []byte
}

func Test1(t *testing.T) {

	params := parameters{s: []byte{'h', 'e', 'l', 'l', 'o'}}

	expected := []byte{'o', 'l', 'l', 'e', 'h'}

	if testResult := reverseString(params.s); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{s: []byte{'H', 'a', 'n', 'n', 'a', 'h'}}

	expected := []byte{'h', 'a', 'n', 'n', 'a', 'H'}

	if testResult := reverseString(params.s); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

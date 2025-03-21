package Q197

import (
	"reflect"
	"testing"
)

type parameters struct {
	s string
	c byte
}

func Test1(t *testing.T) {

	params := parameters{
		s: "loveleetcode",
		c: 'e',
	}

	expected := []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}

	if testResult := shortestToChar(params.s, params.c); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{
		s: "aaab",
		c: 'b',
	}

	expected := []int{3, 2, 1, 0}

	if testResult := shortestToChar(params.s, params.c); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

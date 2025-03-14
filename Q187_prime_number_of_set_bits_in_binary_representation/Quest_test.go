package Q187

import (
	"reflect"
	"testing"
)

type parameters struct {
	left, right int
}

func Test1(t *testing.T) {

	params := parameters{left: 6, right: 10}

	expected := 4

	if testResult := countPrimeSetBits(params.left, params.right); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{left: 10, right: 15}

	expected := 5

	if testResult := countPrimeSetBits(params.left, params.right); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

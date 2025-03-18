package Q194

import (
	"reflect"
	"testing"
)

type parameters struct {
	widths []int
	str    string
}

func Test1(t *testing.T) {

	params := parameters{
		widths: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		str:    "abcdefghijklmnopqrstuvwxyz",
	}

	expected := []int{3, 60}

	if testResult := numberOfLines(params.widths, params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{
		widths: []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		str:    "bbbcccdddaaa",
	}

	expected := []int{2, 4}

	if testResult := numberOfLines(params.widths, params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

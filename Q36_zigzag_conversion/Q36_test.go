package Q36

import (
	"reflect"
	"testing"
)

type parameters struct {
	s       string
	numRows int
}

func Test1(t *testing.T) {

	params := parameters{s: "PAYPALISHIRING", numRows: 3}

	expected := "PAHNAPLSIIGYIR"

	if testResult := convert(params.s, params.numRows); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{s: "PAYPALISHIRING", numRows: 4}

	expected := "PINALSIGYAHRPI"

	if testResult := convert(params.s, params.numRows); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{s: "A", numRows: 1}

	expected := "A"

	if testResult := convert(params.s, params.numRows); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

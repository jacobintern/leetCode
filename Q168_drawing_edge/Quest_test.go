package Q168

import (
	"reflect"
	"testing"
)

type parameters struct {
	n int32
}

func Test1(t *testing.T) {

	params := parameters{n: 2}

	expected := int32(2)

	if testResult := drawingEdge(params.n); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{n: 3}

	expected := int32(8)

	if testResult := drawingEdge(params.n); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{n: 4}

	expected := int32(64)

	if testResult := drawingEdge(params.n); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{n: 911}

	expected := int32(64)

	if testResult := drawingEdge(params.n); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

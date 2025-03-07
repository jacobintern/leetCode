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

	expected := 2

	if testResult := drawingEdge(params.n); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{n: 3}

	expected := 8

	if testResult := drawingEdge(params.n); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{n: 4}

	expected := int64(64)

	if testResult := drawingEdge(params.n); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{n: 911}

	expected := int64(984726690)

	if testResult := drawingEdge(params.n); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test5(t *testing.T) {

	params := parameters{n: 10}

	expected := int64(511620083)

	if testResult := drawingEdge(params.n); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test6(t *testing.T) {

	params := parameters{n: 5}

	expected := int64(1024)

	if testResult := drawingEdge(params.n); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

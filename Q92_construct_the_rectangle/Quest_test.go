package Q92

import (
	"reflect"
	"testing"
)

type parameters struct {
	area int
}

func Test1(t *testing.T) {

	params := parameters{area: 4}

	expected := []int{2, 2}

	if testResult := constructRectangle(params.area); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{area: 37}

	expected := []int{37, 1}

	if testResult := constructRectangle(params.area); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{area: 122122}

	expected := []int{427, 286}

	if testResult := constructRectangle(params.area); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

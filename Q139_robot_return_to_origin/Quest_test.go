package Q139

import (
	"reflect"
	"testing"
)

type parameters struct {
	str string
}

func Test1(t *testing.T) {

	params := parameters{str: "UD"}

	expected := true

	if testResult := judgeCircle(params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{str: "LL"}

	expected := false

	if testResult := judgeCircle(params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

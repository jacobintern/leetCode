package Q202

import (
	"reflect"
	"testing"
)

type parameters struct {
	str, str2 string
}

func Test1(t *testing.T) {

	params := parameters{str: "ab#c", str2: "ad#c"}

	expected := true

	if testResult := backspaceCompare(params.str, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{str: "ab##", str2: "c#d#"}

	expected := true

	if testResult := backspaceCompare(params.str, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{str: "a#c", str2: "b"}

	expected := false

	if testResult := backspaceCompare(params.str, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{str: "xywrrmp", str2: "xywrrmu#p"}

	expected := true

	if testResult := backspaceCompare(params.str, params.str2); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

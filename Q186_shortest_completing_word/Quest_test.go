package Q186

import (
	"reflect"
	"testing"
)

type parameters struct {
	licensePlate string
	words        []string
}

func Test1(t *testing.T) {

	params := parameters{
		licensePlate: "1s3 PSt",
		words:        []string{"step", "steps", "stripe", "stepple"},
	}

	expected := "steps"

	if testResult := shortestCompletingWord(params.licensePlate, params.words); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{
		licensePlate: "1s3 456",
		words:        []string{"looks", "pest", "stew", "show"},
	}

	expected := "pest"

	if testResult := shortestCompletingWord(params.licensePlate, params.words); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

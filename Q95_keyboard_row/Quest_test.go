package Q95

import (
	"reflect"
	"testing"
)

type parameters struct {
	strArr []string
}

func Test1(t *testing.T) {

	params := parameters{strArr: []string{"Hello", "Alaska", "Dad", "Peace"}}

	expected := []string{"Alaska", "Dad"}

	if testResult := findWords(params.strArr); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{strArr: []string{"omk"}}

	expected := []string{}

	if testResult := findWords(params.strArr); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{strArr: []string{"adsdf", "sfd"}}

	expected := []string{"adsdf", "sfd"}

	if testResult := findWords(params.strArr); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

package Q74

import (
	"reflect"
	"testing"
)

type parameters struct {
	ransomNote string
	magazine   string
}

func Test1(t *testing.T) {

	params := parameters{ransomNote: "a", magazine: "b"}

	expected := false

	if testResult := canConstruct(params.ransomNote, params.magazine); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{ransomNote: "aa", magazine: "ab"}

	expected := false

	if testResult := canConstruct(params.ransomNote, params.magazine); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{ransomNote: "aa", magazine: "aab"}

	expected := true

	if testResult := canConstruct(params.ransomNote, params.magazine); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

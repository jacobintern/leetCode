package Q193

import (
	"reflect"
	"testing"
)

type parameters struct {
	words []string
}

func Test1(t *testing.T) {

	params := parameters{words: []string{"gin", "zen", "gig", "msg"}}

	expected := 2

	if testResult := uniqueMorseRepresentations(params.words); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{words: []string{"a"}}

	expected := 1

	if testResult := uniqueMorseRepresentations(params.words); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

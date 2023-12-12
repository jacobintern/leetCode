package Q136

import (
	"reflect"
	"testing"
)

type parameters struct {
	words []string
}

func Test1(t *testing.T) {

	params := parameters{words: []string{"bella", "label", "roller"}}

	expected := []string{"e", "l", "l"}

	if testResult := commonChars(params.words); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{words: []string{"cool", "lock", "cook"}}

	expected := []string{"c", "o"}

	if testResult := commonChars(params.words); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{words: []string{"bbddabab", "cbcddbdd", "bbcadcab", "dabcacad", "cddcacbc", "ccbdbcba", "cbddaccc", "accdcdbb"}}

	expected := []string{"b", "d"}

	if testResult := commonChars(params.words); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

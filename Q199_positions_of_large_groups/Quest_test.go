package Q199

import (
	"reflect"
	"testing"
)

type parameters struct {
	str string
}

func Test1(t *testing.T) {

	params := parameters{str: "abbxxxxzzy"}

	expected := [][]int{{3, 6}}

	if testResult := largeGroupPositions(params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{str: "abc"}

	expected := [][]int{}

	if testResult := largeGroupPositions(params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{str: "abcdddeeeeaabbbcd"}

	expected := [][]int{{3, 5}, {6, 9}, {12, 14}}

	if testResult := largeGroupPositions(params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{str: "aaa"}

	expected := [][]int{{0, 2}}

	if testResult := largeGroupPositions(params.str); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

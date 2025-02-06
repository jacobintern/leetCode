package Q32

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums []int
}

func Test1(t *testing.T) {

	params := parameters{
		nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
	}

	expected := []int{5, 6}

	if testResult := findDisappearedNumbers(params.nums); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{
		nums: []int{1, 1},
	}

	expected := []int{2}

	if testResult := findDisappearedNumbers(params.nums); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

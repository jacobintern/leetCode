package Q97

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums []int
}

func Test1(t *testing.T) {

	params := parameters{nums: []int{5, 4, 3, 2, 1}}

	expected := []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}

	if testResult := findRelativeRanks(params.nums); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{nums: []int{10, 3, 8, 9, 4}}

	expected := []string{"Gold Medal", "4", "Bronze Medal", "Silver Medal", "5"}

	if testResult := findRelativeRanks(params.nums); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

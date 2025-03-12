package Q184

import (
	"reflect"
	"testing"
)

type parameters struct {
	cost []int
}

func Test1(t *testing.T) {

	params := parameters{cost: []int{10, 15, 20}}

	expected := 15

	if testResult := minCostClimbingStairs(params.cost); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}}

	expected := 6

	if testResult := minCostClimbingStairs(params.cost); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

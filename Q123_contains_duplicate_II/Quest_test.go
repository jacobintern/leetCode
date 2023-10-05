package Q123

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums []int
	k    int
}

func Test1(t *testing.T) {

	params := parameters{nums: []int{1, 2, 3, 1}, k: 3}

	expected := true

	if testResult := containsNearbyDuplicate(params.nums, params.k); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{nums: []int{1, 0, 1, 1}, k: 1}

	expected := true

	if testResult := containsNearbyDuplicate(params.nums, params.k); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}
func Test3(t *testing.T) {

	params := parameters{nums: []int{1, 2, 3, 1, 2, 3}, k: 2}

	expected := false

	if testResult := containsNearbyDuplicate(params.nums, params.k); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

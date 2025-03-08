package Q174

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums []int
	adds []int
}

func Test1(t *testing.T) {

	params := parameters{
		nums: []int{4, 5, 8, 2},
		adds: []int{3, 5, 10, 9, 4},
	}

	expectedArr := []int{4, 5, 5, 8, 8}

	kth := Constructor(3, params.nums)

	for i, expected := range expectedArr {
		if testResult := kth.Add(params.adds[i]); reflect.DeepEqual(expected, testResult) {
			t.Log("success")
		} else {
			t.Error("fail coz expected is ", expected, " and test result is ", testResult)
		}
	}
}

func Test2(t *testing.T) {

	params := parameters{
		nums: []int{7, 7, 7, 7, 8, 3},
		adds: []int{2, 10, 9, 9},
	}

	expectedArr := []int{7, 7, 7, 8}

	kth := Constructor(4, params.nums)

	for i, expected := range expectedArr {
		if testResult := kth.Add(params.adds[i]); reflect.DeepEqual(expected, testResult) {
			t.Log("success")
		} else {
			t.Error("fail coz expected is ", expected, " and test result is ", testResult)
		}
	}
}

package Q40

import (
	"reflect"
	"testing"
)

type parameters struct {
	arr []int
}

func Test1(t *testing.T) {

	params := parameters{arr: []int{-1, 0, 1, 2, -1, -4}}

	expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}

	if testResult := threeSum(params.arr); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

// func Test2(t *testing.T) {

// 	params := parameters{arr: []int{}}

// 	expected := 0

// 	if testResult := threeSum(params.arr); reflect.DeepEqual(expected, testResult) {
// 		t.Log("sucess")
// 	} else {
// 		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
// 	}
// }

// func Test3(t *testing.T) {

// 	params := parameters{arr: []int{-1, 0, 1, 2, -1, -4}}

// 	expected := 0

// 	if testResult := threeSum(params.arr); reflect.DeepEqual(expected, testResult) {
// 		t.Log("sucess")
// 	} else {
// 		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
// 	}
// }

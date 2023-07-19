package Q45

import (
	"reflect"
	"testing"
)

type parameters struct {
	arr []int
}

func Test1(t *testing.T) {

	params := parameters{arr: []int{0, 1, 2, 4, 5, 7}}

	expected := []string{"0->2", "4->5", "7"}

	if testResult := summaryRanges(params.arr); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{arr: []int{0, 2, 3, 4, 6, 8, 9}}

	expected := []string{"0", "2->4", "6", "8->9"}

	if testResult := summaryRanges(params.arr); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

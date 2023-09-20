package Q113

import (
	"reflect"
	"testing"
)

type parameters struct {
	arr []int
}

func Test1(t *testing.T) {

	params := parameters{arr: []int{1, 3, 2, 2, 5, 2, 3, 7}}

	expected := 5

	if testResult := findLHS(params.arr); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

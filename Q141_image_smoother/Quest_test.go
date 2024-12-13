package Q141

import (
	"reflect"
	"testing"
)

type parameters struct {
	arrList [][]int
}

func Test1(t *testing.T) {

	params := parameters{arrList: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}}

	expected := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	if testResult := imageSmoother(params.arrList); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{arrList: [][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}}

	expected := [][]int{{137, 141, 137}, {141, 138, 141}, {137, 141, 137}}

	if testResult := imageSmoother(params.arrList); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{arrList: [][]int{{2, 3, 4}, {5, 6, 7}, {8, 9, 10}, {11, 12, 13}, {14, 15, 16}}}

	expected := [][]int{{4, 4, 5}, {5, 6, 6}, {8, 9, 9}, {11, 12, 12}, {13, 13, 14}}

	if testResult := imageSmoother(params.arrList); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

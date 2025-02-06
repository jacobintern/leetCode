package Q26

import (
	"reflect"
	"testing"
)

type parameters struct {
	numRows int
}

func Test1(t *testing.T) {
	params := parameters{
		numRows: 5,
	}
	expected := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}

	if reflect.DeepEqual(expected, generate(params.numRows)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		numRows: 1,
	}
	expected := [][]int{{1}}

	if reflect.DeepEqual(expected, generate(params.numRows)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

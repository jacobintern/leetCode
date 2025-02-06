package Q27

import (
	"reflect"
	"testing"
)

type parameters struct {
	rowIndex int
}

func Test1(t *testing.T) {
	params := parameters{
		rowIndex: 5,
	}
	expected := []int{1, 4, 6, 4, 1}

	if reflect.DeepEqual(expected, getRow(params.rowIndex)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		rowIndex: 1,
	}
	expected := []int{1}

	if reflect.DeepEqual(expected, getRow(params.rowIndex)) {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}

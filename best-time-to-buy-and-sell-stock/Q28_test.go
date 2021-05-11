package Q28

import (
	"reflect"
	"testing"
)

type parameters struct {
	prices []int
}

func Test1(t *testing.T) {
	params := parameters{
		prices: []int{7, 1, 5, 3, 6, 4},
	}
	expected := 5

	if reflect.DeepEqual(expected, maxProfit(params.prices)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {
	params := parameters{
		prices: []int{7, 6, 4, 3, 1},
	}
	expected := 0

	if reflect.DeepEqual(expected, maxProfit(params.prices)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

package Q148

import (
	"reflect"
	"testing"
)

type parameters struct {
	customers []int
	grumpy    []int
	minutes   int
}

func Test1(t *testing.T) {

	params := parameters{
		customers: []int{1, 0, 1, 2, 1, 1, 7, 5},
		grumpy:    []int{0, 1, 0, 1, 0, 1, 0, 1},
		minutes:   3,
	}

	expected := 16

	if testResult := maxSatisfied(params.customers, params.grumpy, params.minutes); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{
		customers: []int{1},
		grumpy:    []int{0},
		minutes:   1,
	}

	expected := 1

	if testResult := maxSatisfied(params.customers, params.grumpy, params.minutes); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{
		customers: []int{5, 8},
		grumpy:    []int{0, 1},
		minutes:   1,
	}

	expected := 13

	if testResult := maxSatisfied(params.customers, params.grumpy, params.minutes); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{
		customers: []int{4, 10, 10},
		grumpy:    []int{1, 1, 0},
		minutes:   2,
	}

	expected := 24

	if testResult := maxSatisfied(params.customers, params.grumpy, params.minutes); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test5(t *testing.T) {

	params := parameters{
		customers: []int{7, 8, 8, 6},
		grumpy:    []int{0, 1, 0, 1},
		minutes:   3,
	}

	expected := 29

	if testResult := maxSatisfied(params.customers, params.grumpy, params.minutes); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test6(t *testing.T) {

	params := parameters{
		customers: []int{3, 8, 8, 7, 1},
		grumpy:    []int{1, 1, 1, 1, 1},
		minutes:   3,
	}

	expected := 23

	if testResult := maxSatisfied(params.customers, params.grumpy, params.minutes); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

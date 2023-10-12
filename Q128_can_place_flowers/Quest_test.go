package Q128

import (
	"reflect"
	"testing"
)

type parameters struct {
	flowerbed []int
	n         int
}

func Test1(t *testing.T) {

	params := parameters{flowerbed: []int{1, 0, 0, 0, 1}, n: 1}

	expected := true

	if testResult := canPlaceFlowers(params.flowerbed, params.n); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{flowerbed: []int{1, 0, 0, 0, 1}, n: 2}

	expected := false

	if testResult := canPlaceFlowers(params.flowerbed, params.n); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{flowerbed: []int{0, 1, 0}, n: 1}

	expected := false

	if testResult := canPlaceFlowers(params.flowerbed, params.n); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

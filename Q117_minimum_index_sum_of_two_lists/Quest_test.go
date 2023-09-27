package Q117

import (
	"reflect"
	"testing"
)

type parameters struct {
	list1 []string
	list2 []string
}

func Test1(t *testing.T) {

	params := parameters{list1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		list2: []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}}

	expected := []string{"Shogun"}

	if testResult := findRestaurant(params.list1, params.list2); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{list1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		list2: []string{"KFC", "Shogun", "Burger King"}}

	expected := []string{"Shogun"}

	if testResult := findRestaurant(params.list1, params.list2); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{list1: []string{"happy", "sad", "good"},
		list2: []string{"sad", "happy", "good"}}

	expected := []string{"sad", "happy"}

	if testResult := findRestaurant(params.list1, params.list2); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {

	params := parameters{list1: []string{"vacag", "KFC"},
		list2: []string{"fvo", "xrljq", "jrl", "KFC"}}

	expected := []string{"KFC"}

	if testResult := findRestaurant2(params.list1, params.list2); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

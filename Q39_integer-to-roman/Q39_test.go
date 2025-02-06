package Q39

import (
	"reflect"
	"testing"
)

type parameters struct {
	num int
}

func Test1(t *testing.T) {
	params := parameters{
		num: 3,
	}
	expected := "III"

	if testResult := intToRoman(params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {
	params := parameters{
		num: 4,
	}
	expected := "IV"

	if testResult := intToRoman(params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {
	params := parameters{
		num: 9,
	}
	expected := "IX"

	if testResult := intToRoman(params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test4(t *testing.T) {
	params := parameters{
		num: 58,
	}
	expected := "LVIII"

	if testResult := intToRoman(params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test5(t *testing.T) {
	params := parameters{
		num: 1994,
	}
	expected := "MCMXCIV"

	if testResult := intToRoman(params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

package Q189

import (
	"reflect"
	"testing"
)

type parameters struct {
	jewels string
	stones string
}

func Test1(t *testing.T) {

	params := parameters{jewels: "aA", stones: "aAAbbbb"}

	expected := 3

	if testResult := numJewelsInStones(params.jewels, params.stones); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{jewels: "z", stones: "ZZ"}

	expected := 0

	if testResult := numJewelsInStones(params.jewels, params.stones); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

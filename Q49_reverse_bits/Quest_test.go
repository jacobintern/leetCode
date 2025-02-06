package Q49

import (
	"reflect"
	"testing"
)

type parameters struct {
	num uint32
}

func Test1(t *testing.T) {

	params := parameters{num: 43261596}

	expected := uint32(964176192)

	if testResult := reverseBits(params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{num: 4294967293}

	expected := uint32(3221225471)

	if testResult := reverseBits2(params.num); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

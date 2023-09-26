package Q30

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {

	str := "A man, a plan, a canal: Panama"

	expected := true

	if reflect.DeepEqual(expected, isPalindrome3(str)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test2(t *testing.T) {

	str := "race a car"

	expected := false

	if reflect.DeepEqual(expected, isPalindrome2(str)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test3(t *testing.T) {

	str := " "

	expected := true

	if reflect.DeepEqual(expected, isPalindrome(str)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

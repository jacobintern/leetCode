package Q100

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {

	expected := 0

	if testResult := Quest(); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

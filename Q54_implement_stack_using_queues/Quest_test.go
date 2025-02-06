package Q54

import (
	"reflect"
	"testing"
)

type parameters struct {
	arr     []string
	pushArr []int
}

func Test1(t *testing.T) {

	params := parameters{
		arr:     []string{"MyStack", "push", "push", "top", "pop", "empty"},
		pushArr: []int{0, 1, 2, 0, 0, 0},
	}
	obj := Constructor()

	for i, v := range params.arr {
		switch v {
		case "MyStack":
			expected := []int{}
			if testResult := obj.stack; reflect.DeepEqual(expected, testResult) {
				t.Log("success")
			} else {
				t.Error("fail coz expected is ", expected, " and test result is ", testResult)
			}
		case "push":
			obj.Push(params.pushArr[i])
		case "top":
			expected := 2
			if testResult := obj.Top(); reflect.DeepEqual(expected, testResult) {
				t.Log("success")
			} else {
				t.Error("fail coz expected is ", expected, " and test result is ", testResult)
			}
		case "pop":
			expected := 2
			if testResult := obj.Pop(); reflect.DeepEqual(expected, testResult) {
				t.Log("success")
			} else {
				t.Error("fail coz expected is ", expected, " and test result is ", testResult)
			}
		case "empty":
			expected := false
			if testResult := obj.Empty(); reflect.DeepEqual(expected, testResult) {
				t.Log("success")
			} else {
				t.Error("fail coz expected is ", expected, " and test result is ", testResult)
			}
		}
	}
}

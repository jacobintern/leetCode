package Q56

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
		arr:     []string{"MyQueue", "push", "push", "peek", "pop", "empty"},
		pushArr: []int{0, 1, 2, 0, 0, 0},
	}
	obj := Constructor()

	for i, v := range params.arr {
		switch v {
		case "MyQueue":
			expected := []int{}
			if testResult := obj.stack; reflect.DeepEqual(expected, testResult) {
				t.Log("sucess")
			} else {
				t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
			}
		case "push":
			obj.Push(params.pushArr[i])
		case "peek":
			expected := 1
			if testResult := obj.Peek(); reflect.DeepEqual(expected, testResult) {
				t.Log("sucess")
			} else {
				t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
			}
		case "pop":
			expected := 1
			if testResult := obj.Pop(); reflect.DeepEqual(expected, testResult) {
				t.Log("sucess")
			} else {
				t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
			}
		case "empty":
			expected := false
			if testResult := obj.Empty(); reflect.DeepEqual(expected, testResult) {
				t.Log("sucess")
			} else {
				t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
			}
		}
	}
}

func Test2(t *testing.T) {

	params := parameters{
		arr:     []string{"MyQueue", "push", "pop", "empty"},
		pushArr: []int{0, 1, 0, 0},
	}
	obj := Constructor()

	for i, v := range params.arr {
		switch v {
		case "MyQueue":
			expected := []int{}
			if testResult := obj.stack; reflect.DeepEqual(expected, testResult) {
				t.Log("sucess")
			} else {
				t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
			}
		case "push":
			obj.Push(params.pushArr[i])
		case "peek":
			expected := 1
			if testResult := obj.Peek(); reflect.DeepEqual(expected, testResult) {
				t.Log("sucess")
			} else {
				t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
			}
		case "pop":
			expected := 1
			if testResult := obj.Pop(); reflect.DeepEqual(expected, testResult) {
				t.Log("sucess")
			} else {
				t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
			}
		case "empty":
			expected := false
			if testResult := obj.Empty(); reflect.DeepEqual(expected, testResult) {
				t.Log("sucess")
			} else {
				t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
			}
		}
	}
}

package Q176

import (
	"reflect"
	"testing"
)

type parameters struct {
	actions []func() any
}

func Test1(t *testing.T) {
	heap := Constructor()

	// params := parameters{
	actions := []func() any{
		func() any { return nil },
		func() any { heap.Add(1); return nil },
		func() any { heap.Add(2); return nil },
		func() any { return heap.Contains(1) },
		func() any { return heap.Contains(3) },
		func() any { heap.Add(2); return nil },
		func() any { return heap.Contains(2) },
		func() any { heap.Remove(2); return nil },
		func() any { return heap.Contains(2) },
	}

	expected := []any{nil, nil, nil, true, false, nil, true, nil, false}

	for i, action := range actions {
		if testResult := action(); reflect.DeepEqual(expected[i], testResult) {
			t.Log("success")
		} else {
			t.Error("fail coz expected is ", expected[i], " and test result is ", testResult)
		}
	}
}

func Test2(t *testing.T) {
	heap := Constructor()

	// params := parameters{
	actions := []func() any{
		func() any { return nil },
		func() any { heap.Add(9); return nil },
		func() any { heap.Remove(19); return nil },
		func() any { heap.Add(14); return nil },
		func() any { heap.Remove(19); return nil },
		func() any { heap.Remove(9); return nil },
		func() any { heap.Add(0); return nil },
		func() any { heap.Add(3); return nil },
		func() any { heap.Add(4); return nil },
		func() any { heap.Add(0); return nil },
		func() any { heap.Remove(9); return nil },
	}

	expected := []any{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}

	for i, action := range actions {
		if testResult := action(); reflect.DeepEqual(expected[i], testResult) {
			t.Log("success")
		} else {
			t.Error("fail coz expected is ", expected[i], " and test result is ", testResult)
		}
	}
}

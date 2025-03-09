package Q177

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
		func() any { heap.Put(1, 1); return nil },
		func() any { heap.Put(2, 2); return nil },
		func() any { return heap.Get(1) },
		func() any { return heap.Get(3) },
		func() any { heap.Put(2, 1); return nil },
		func() any { return heap.Get(2) },
		func() any { heap.Remove(2); return nil },
		func() any { return heap.Get(2) },
	}

	expected := []any{nil, nil, nil, 1, -1, nil, 1, nil, -1}

	for i, action := range actions {
		if testResult := action(); reflect.DeepEqual(expected[i], testResult) {
			t.Log("success")
		} else {
			t.Error("fail coz expected is ", expected[i], " and test result is ", testResult)
		}
	}
}

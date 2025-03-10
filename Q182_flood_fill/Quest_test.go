package Q182

import (
	"reflect"
	"testing"
)

type parameters struct {
	image         [][]int
	sr, sc, color int
}

func Test1(t *testing.T) {

	params := parameters{
		image: [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
		sr:    1,
		sc:    1,
		color: 2,
	}

	expected := [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}

	if testResult := floodFill(params.image, params.sr, params.sc, params.color); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

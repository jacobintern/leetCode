package Q31

import (
	"reflect"
	"testing"
)

type parameters struct {
	nums []int
}

func Test1(t *testing.T) {

	params := parameters{
		nums: []int{2, 2, 1, 11, 1, 1123, 1123},
	}

	expected := 11

	if testResult := singleNumber(params.nums); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{
		nums: []int{4, 1, 2, 1, 2},
	}

	expected := 4

	if reflect.DeepEqual(expected, singleNumber(params.nums)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test3(t *testing.T) {

	params := parameters{
		nums: []int{1},
	}

	expected := 1

	if reflect.DeepEqual(expected, singleNumber(params.nums)) {
		t.Log("sucess")
	} else {
		t.Error("fail")
	}
}

func Test4(t *testing.T) {

	params := parameters{
		nums: []int{463, 191, -915, -438, 420, -135, -431, 611, 695, -605, 469, 917, 301, 70, 209, -66, 956, 7, 245, 147, 104, 633, -218, -879, -894, 208, -37, -344, -252, -684, 728, -943, 858, -557, 217, -655, -91, -308, 699, -152, -702, 916, -867, 447, 29, -207, 56, -149, 909, 980, 508, -747, -389, -718, 814, -790, 803, 299, 443, 932, -814, 495, 274, 88, -22, 373, -324, 144, -144, 421, -499, -178, -582, -124, 741, 526, 215, -538, -912, 400, 425, -693, -767, 862, -746, 462, 762, 148, 463, 191, -915, -438, 420, -135, -431, 611, 695, -605, 469, 917, 301, 70, 209, -66, 956, 7, 245, 147, 104, 633, -218, -879, -894, 208, -37, -344, -252, -684, 728, -943, 858, -557, 217, -655, -91, -308, 699, -152, -702, 916, -867, 447, 29, -207, 56, -149, 909, 980, 508, -747, -389, -718, 814, -790, 803, 299, 443, 932, -814, 495, 274, 88, -22, 373, -324, 144, -144, 421, -499, -178, -582, -124, 741, 526, 215, -538, -912, 400, 425, -693, -767, 862, -746, 462, 762, 148, -3},
	}

	expected := -3

	if testResult := singleNumber(params.nums); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

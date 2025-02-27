package Q162

import (
	"reflect"
	"testing"
)

type parameters struct {
	strs []string
}

func Test1(t *testing.T) {

	params := parameters{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}}

	expected := [][]string{{"bat"}, {"tan", "nat"}, {"eat", "ate", "tea"}}

	if testResult := groupAnagrams(params.strs); reflect.DeepEqual(expected, testResult) {
		t.Log("success")
	} else {
		t.Error("fail coz expected is ", expected, " and test result is ", testResult)
	}
}

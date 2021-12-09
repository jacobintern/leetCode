package error_hanlding

import (
	"fmt"
	"testing"
)

type parameters struct {
	str string
}

func Test1(t *testing.T) {

	params := parameters{str: "Jacob"}

	if _, err := checkMember(params.str); err != nil {
		fmt.Println(err)
	}
}

func Test2(t *testing.T) {

	params := parameters{str: "doggg"}

	if _, err := checkMember(params.str); err != nil {
		fmt.Println(err)
	}
}

func Test3(t *testing.T) {

	params := parameters{str: "doggy"}

	if _, err := checkMember(params.str); err != nil {
		fmt.Println(err)
	}
}

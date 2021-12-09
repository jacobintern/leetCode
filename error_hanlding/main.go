package error_hanlding

import (
	"fmt"
)

var userList []string = []string{
	"Jacob",
	"Judy",
	"Bear",
	"Dragon",
	"Giraffa",
	"Koala",
	"Doggy",
}

func checkMember(usrName string) (bool, error) {
	for _, usr := range userList {
		if usr == usrName {
			return true, fmt.Errorf("the user %s is already exist! ", usrName)
		}
	}
	return false, nil
}

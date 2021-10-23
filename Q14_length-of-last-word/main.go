package Q14

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	strArr := strings.Split(s, " ")
	arrLen := len(strArr)
	r := ""
	for arrLen > 0 {
		if strArr[arrLen-1] != "" {
			r = strArr[arrLen-1]
			break
		}
		arrLen--
	}
	return len(r)
}

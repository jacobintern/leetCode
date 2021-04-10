package Q3

import (
	"strconv"
	"strings"
)

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strings.Split(strconv.Itoa(x), "")
	for i, v := range s {
		if s[(len(s)-1)-i] != v {
			return false
		}
	}

	return true
}

func IsPalindrome2(x int) bool {
	vStr := strconv.Itoa(x)
	s := strings.Split(vStr, "")
	var newStr string

	i := 0
	for i < len(s) {
		newStr += s[len(s)-1-i]
		i++
	}
	return vStr == newStr
}

package Q30

import (
	"log"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {

	reg, err := regexp.Compile("[^a-zA-Z]+")

	if err != nil {
		log.Fatal(err)
	}

	s = strings.ReplaceAll(s, " ", "")
	s = reg.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	sLen := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[sLen-i] {
			return false
		}
	}
	return true
}

func isPalindrome2(s string) bool {

	s = strings.ToLower(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for ; i < j && !isLetter(s[i]); i++ {
		}
		for ; i < j && !isLetter(s[j]); j-- {
		}
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func isLetter(s byte) bool {
	return s >= 'a' && s <= 'z' || s >= '0' && s <= '9'
}

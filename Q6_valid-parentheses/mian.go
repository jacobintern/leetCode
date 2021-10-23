package Q6

import (
	"strings"
)

func isValid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	if len(s)%2 != 0 {
		return false
	} else {
		l := len(s)
		for l > 0 {
			s = strings.ReplaceAll(s, "()", "")
			s = strings.ReplaceAll(s, "[]", "")
			s = strings.ReplaceAll(s, "{}", "")
			l--
		}
	}
	if s == "" {
		return true
	}
	return false
}

package Q34

import (
	"strings"
)

var curMax int

func lengthOfLongestSubstring(s string) int {
	curMax = 0
	CountLen(s)
	return curMax
}

func CountLen(s string) {
	if len(s) == 0 {
		return
	}
	tmpMax := 0
	tmp := ""
	for _, str := range s {
		if !strings.Contains(tmp, string(str)) {
			tmp += string(str)
			tmpMax++
		} else {
			break
		}
	}
	if tmpMax > curMax {
		curMax = tmpMax
	}
	CountLen(s[1:])
}

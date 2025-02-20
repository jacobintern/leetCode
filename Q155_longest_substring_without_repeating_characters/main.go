package Q156

import "strings"

func lengthOfLongestSubstring(s string) int {
	start, end, str, max := 0, 0, "", 0

	for start < len(s) && end < len(s) {
		tmp := string(s[end])
		if strings.Contains(str, tmp) {
			if l := len(str); l > max {
				max = len(str)
			}
			str, start, end = "", start+1, start+1
			continue
		}

		str += tmp
		end++
	}

	if l := len(str); l > max {
		return l
	}

	return max
}

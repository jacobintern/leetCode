package Q10

import (
	"strings"
)

func strStr(haystack string, needle string) int {
	if haystack == "" && needle == "" {
		return 0
	} else if haystack == "" {
		return -1
	} else if needle == "" {
		return 0
	} else if haystack == needle {
		return 0
	}
	return strings.Index(haystack, needle)
}

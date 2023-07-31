package Q84

import "strings"

func countSegments(s string) int {
	count := 0
	for _, v := range strings.Split(s, " ") {
		if v != "" {
			count++
		}
	}
	return count
}

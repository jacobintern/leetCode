package Q63

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	m := make(map[byte]string)
	m2 := make(map[string]byte)
	strArr := strings.Split(s, " ")
	if len(pattern) != len(strArr) {
		return false
	}

	for i, v := range strArr {
		v1, ok1 := m[pattern[i]]
		v2, ok2 := m2[v]

		if (ok1 && v1 != v) || (ok2 && v2 != pattern[i]) {
			return false
		} else if v1 == v {
			continue
		} else {
			m[pattern[i]] = v
			m2[v] = pattern[i]
		}
	}

	return true
}

package Q136

import "strings"

func commonChars(words []string) []string {
	m := make(map[string]int)
	res := []string{}

	for _, str := range words {
		tmp := make(map[string]int)
		for _, s := range str {
			tmp[string(s)]++
		}

		if str == words[0] {
			m = tmp
		} else {
			// check min
			for i, v := range tmp {
				value, exists := m[i]
				if exists && value > v {
					m[i] = v
				}
			}
		}
	}

	for i := range m {
		if !charExists(i, words) {
			m[i] = 0
		}
	}

	for i, v := range m {
		for v > 0 {
			res = append(res, i)
			v--
		}
	}

	return res
}

func charExists(s string, slice []string) bool {
	for _, str := range slice {
		if !strings.Contains(str, s) {
			return false
		}
	}
	return true
}

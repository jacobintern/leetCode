package Q74

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, v := range ransomNote {
		m[v]++
	}

	for _, v := range magazine {
		count, ok := m[v]
		if ok && count > 0 {
			m[v]--
			if m[v] == 0 {
				delete(m, v)
			}
		}
	}

	if len(m) > 0 {
		return false
	}

	return true
}

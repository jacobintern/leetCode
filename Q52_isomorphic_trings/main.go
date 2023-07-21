package Q52

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	smap := make(map[byte]int)
	tmap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		smap[s[i]]++
		tmap[t[i]]++
		if smap[s[i]] == tmap[t[i]] && i != 0 && s[i] == t[i] && s[i-1] == t[i-1] {
			continue
		} else if smap[s[i]] == tmap[t[i]] && i != 0 && s[i-1] == t[i-1] && s[i] != t[i] {
			return false
		} else if smap[s[i]] == tmap[t[i]] && s[i] == t[i] && i != 0 && (s[i] == s[i-1] || t[i] == t[i-1]) {
			return false
		} else if smap[s[i]] == tmap[t[i]] && s[i] != t[i] && i != 0 && ((s[i] != s[i-1] && t[i] == t[i-1]) || (s[i] == s[i-1] && t[i] != t[i-1])) {
			return false
		} else if smap[s[i]] != tmap[t[i]] {
			return false
		}
	}

	return true
}

func isIsomorphic2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if val, ok := sMap[s[i]]; ok {
			if val != t[i] {
				return false
			}
		} else {
			sMap[s[i]] = t[i]
		}

		if val, ok := tMap[t[i]]; ok {
			if val != s[i] {
				return false
			}
		} else {
			tMap[t[i]] = s[i]
		}
	}

	return true
}

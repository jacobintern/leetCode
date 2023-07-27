package Q77

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	tmp := 0
	for _, v := range t {
		if tmp < len(s) && byte(v) == s[tmp] {
			tmp++
		}
	}
	return len(s)-tmp == 0
}

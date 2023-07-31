package Q87

func repeatedSubstringPattern(s string) bool {
	if len(s)%2 == 0 && s[:len(s)/2] == s[len(s)/2:] {
		return true
	}
	return false
}

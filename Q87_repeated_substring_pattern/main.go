package Q87

func repeatedSubstringPattern(s string) bool {
	if len(s)%2 == 0 && s[:len(s)/2] == s[len(s)/2:] {
		return true
	}

	// need valid distinct result is one
	return false
}

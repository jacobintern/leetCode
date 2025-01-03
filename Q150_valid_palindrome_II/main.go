package Q150

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for r > l {
		if s[l] != s[r] {
			return isPalindrome(s, l+1, r) || isPalindrome(s, l, r-1)
		} else {
			l++
			r--
		}
	}

	return true
}

func isPalindrome(s string, l, r int) bool {
	for r > l {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

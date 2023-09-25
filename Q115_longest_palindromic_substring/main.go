package Q115

func longestPalindrome(s string) string {
	strLen := len(s)
	if strLen == 0 {
		return ""
	}

	// l 前指標, r 後指標, head & tail 標記最長結果
	var l, r, head, tail int
	for r < strLen {
		if (strLen-r)*2 < (tail-head)+1 {
			break
		}
		for r+1 < strLen && s[l] == s[r+1] {
			r++
		}
		for l-1 >= 0 && r+1 < strLen && s[l-1] == s[r+1] {
			l--
			r++
		}

		if r-l > tail-head {
			head, tail = l, r
		}
		l = (l+r)/2 + 1
		r = l
	}

	return s[head : tail+1]
}

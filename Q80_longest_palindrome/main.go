package Q80

func longestPalindrome(s string) int {
	m := make(map[rune]int)

	for _, v := range s {
		m[v]++
	}

	res := 0

	for _, v := range m {
		if v%2 == 0 {
			res += v
		} else {
			res += v - 1
		}
	}

	if res == 0 && len(s) == 1 {
		return 1
	} else if len(s) > res {
		return res + 1
	}

	return res
}

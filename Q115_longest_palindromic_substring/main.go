package Q115

import (
	"strings"
)

func longestPalindrome(s string) string {
	res, strLen := "", len(s)
	if strLen == 0 {
		return res
	}
	if strLen == 1 {
		return s
	}
	if s == strings.Repeat(string(s[0]), strLen) {
		return s
	}

	// 圈數
	for i := 0; i < strLen-1; i++ {
		// l 前指標, r 後指標
		var tmp string
		var l, r int

		if s[i] != s[i+1] {
			// 中心不一樣時
			tmp, l, r = string(s[i]), i, i
		} else {
			tmp, l, r = "", i, i+1
		}

		for l != -1 || r != strLen {
			if r == strLen || l == -1 || s[l] != s[r] {
				break
			} else if s[l] == s[r] && l != r {
				tmp = string(s[l]) + tmp + string(s[r])
			}
			if l > -1 {
				l--
			}
			if r < strLen {
				r++
			}
		}

		if len(tmp) > 0 && len(tmp) > len(res) {
			res = tmp
		}
	}

	return res
}

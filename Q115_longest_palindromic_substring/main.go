package Q115

import (
	"strings"
)

func longestPalindrome(s string) string {
	res := ""
	if len(s) == 0 {
		return res
	}
	if len(s) == 1 {
		return s
	}
	if s == strings.Repeat(string(s[0]), len(s)) {
		return s
	}

	strLen := len(s)

	// 圈數
	for i := 0; i < strLen; i++ {
		// j 前指標, k 後指標
		tmp, k, l := "", i, i
		for j := 0; j < strLen; j++ {
			if s[k] == s[l] {
				if k != l {
					tmp = string(s[k]) + tmp + string(s[l])
				} else {
					tmp = string(s[k]) + tmp
				}
			}
			if k != 0 {
				k--
			}
			if l != strLen-1 {
				l++
			}
		}
		if len(tmp) > 0 && len(tmp) > len(res) {
			res = tmp
		}
	}

	return res
}

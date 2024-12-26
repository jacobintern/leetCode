package Q143

import (
	"strings"
	"unicode"
)

func licenseKeyFormatting(s string, k int) string {
	l := len(s) - strings.Count(s, "-")
	v, i := l%k, 1
	var res strings.Builder
	for _, str := range s {
		if str == rune('-') {
			continue
		}

		res.WriteRune(unicode.ToUpper(str))

		if (i == v || (i-v)%k == 0) && i != l {
			res.WriteRune('-')
		}

		i++
	}

	return res.String()
}

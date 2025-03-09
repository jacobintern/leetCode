package Q178

import (
	"strings"
)

func toLowerCase(s string) string {
	var res strings.Builder
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			res.WriteRune(v + 32)
		} else {
			res.WriteRune(v)
		}
	}

	return res.String()
}

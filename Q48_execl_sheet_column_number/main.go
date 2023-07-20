package Q48

import (
	"math"
)

func titleToNumber(columnTitle string) int {
	res := 0.0
	for i, j := len(columnTitle), 0; i > 0; i-- {
		res += math.Pow(26, float64(i-1)) * float64(columnTitle[j]-64)
		j++
	}
	return int(res)
}

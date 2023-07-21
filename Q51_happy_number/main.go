package Q51

import (
	"math"
	"strconv"
)

func isHappy(n int) bool {
	tmplist := make(map[int]int)

	for {
		if n == 1 {
			return true
		}

		tmplist[n]++

		if tmplist[n] == 2 {
			return false
		}
		tmp := 0.0
		for i := len(strconv.Itoa(n)); i > 0; i-- {
			tmp += math.Pow(float64(n%10), 2)
			n /= 10
		}
		n = int(tmp)
	}
}

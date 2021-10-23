package Q12

import "math"

func myPow(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}
	if x == -1 && n%2 != 0 {
		return -1
	}
	if x == -1 && n%2 == 0 {
		return 1
	}
	tmp := x
	x = 1
	numlen := math.Abs(float64(n))
	for numlen > 0 {
		x = x * tmp
		numlen--
	}

	if n > 0 {
		return 1 * x
	}
	return 1 / x
}

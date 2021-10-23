package Q17

import "math"

func mySqrt(x int) int {
	res := x
	for res*res > x {
		res = (res + x/res) / 2
	}
	return res
}

func mySqrt2(x int) int {
	return int(math.Sqrt(float64(x)))
}

package Q72

import "math"

func isPerfectSquare(num int) bool {
	v := math.Sqrt(float64(num))

	if v != float64(int(v)) {
		return false
	}

	return true
}

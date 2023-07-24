package Q55

import "math"

func isPowerOfTwo(n int) bool {
	res := math.Log2(float64(n))

	if res != float64(int64(res)) {
		return false
	}

	return true
}

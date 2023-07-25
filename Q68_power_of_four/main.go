package Q68

func isPowerOfFour(n int) bool {
	return isPowerOfN(n, 4)
}

func isPowerOfN(n int, t float64) bool {
	tmp := float64(n)
	if n < 0 || n == 0 {
		return false
	}

	for tmp > 1 {
		tmp /= t
	}

	if tmp == 1 {
		return true
	}

	return false
}

package Q66

func isPowerOfThree(n int) bool {
	tmp := float64(n)
	if n < 0 || n == 0 {
		return false
	}

	for tmp > 1 {
		tmp /= 3
	}

	if tmp == 1 {
		return true
	}

	return false
}

package Q85

func arrangeCoins(n int) int {
	i := 0
	for n >= i {
		i++
		n -= i
	}

	if n < 0 {
		return i - 1
	}

	return i
}

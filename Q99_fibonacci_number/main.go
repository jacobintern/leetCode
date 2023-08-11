package Q99

func fib(n int) int {
	if n < 2 {
		return n
	}

	pre, cur := 0, 1

	for n-1 > 0 {
		pre, cur = cur, pre+cur
		n--
	}

	return cur
}

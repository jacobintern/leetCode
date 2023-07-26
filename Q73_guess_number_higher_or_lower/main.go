package Q73

var target int

func guessNumber(n int) int {
	val := 0
	for {
		val = guess(n)
		if val == 0 {
			return n
		}
		n += val
	}
}

func guess(num int) int {
	if num > target {
		return -1
	} else if num < target {
		return 1
	}

	return 0
}

func setTaget(num int) {
	target = num
}

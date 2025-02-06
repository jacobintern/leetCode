package Q152

func hasAlternatingBits(n int) bool {
	var pre string

	for n > 0 {
		var cur string
		if n%2 == 1 {
			cur = "1"
		} else {
			cur = "0"
		}

		if pre == cur {
			return false
		}

		n = n / 2
		pre = cur
	}

	return true
}

package Q90

func findComplement(num int) int {
	res, i := 0, 1

	for num > 0 {
		if num%2 == 0 {
			res += i
		}
		num /= 2
		i *= 2
	}

	return res
}

package Q98

func checkPerfectNumber(num int) bool {
	sum := 0
	if num%2 != 0 {
		return false
	}

	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			sum += i
		}
	}

	return num == sum
}

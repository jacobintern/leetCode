package Q15

func plusOne(digits []int) []int {
	digitsLen := len(digits) - 1
	digits[digitsLen] += 1
	for digitsLen >= 0 {
		if digits[digitsLen] >= 10 && digitsLen != 0 {
			digits[digitsLen] = 0
			digits[digitsLen-1] += 1
		}
		if digits[digitsLen] >= 10 && digitsLen == 0 {
			digits[digitsLen] = 0
			digits = append([]int{1}, digits...)
		}
		digitsLen--
	}
	return digits
}

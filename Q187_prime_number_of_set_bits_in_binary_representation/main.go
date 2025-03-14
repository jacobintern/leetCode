package Q187

import "math/bits"

func countPrimeSetBits(left int, right int) int {
	count := 0
	for num := left; num <= right; num++ {
		ones := bits.OnesCount(uint(num)) // 計算二進位中 1 的個數
		if isPrime(ones) {
			count++
		}
	}
	return count
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

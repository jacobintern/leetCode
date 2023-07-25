package Q61

import "math/rand"

var _ultimatePassword int32

func init() {
	_ultimatePassword = generateBadVersion()
}

func firstBadVersion(n int) int {
	left := 1
	right := n

	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func isBadVersion(num int) bool {
	return num == int(_ultimatePassword)
}

func generateBadVersion() int32 {
	return rand.Int31()
}

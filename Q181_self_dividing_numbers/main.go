package Q181

import "strconv"

func selfDividingNumbers(left int, right int) []int {
	res := []int{}

	for left <= right {
		if isSelfDividing(left) {
			res = append(res, left)
		}
		left++
	}
	return res
}

func isSelfDividing(n int) bool {
	str := strconv.Itoa(n)
	for _, v := range str {
		if v == '0' || n%int(v-'0') != 0 {
			return false
		}
	}
	return true
}

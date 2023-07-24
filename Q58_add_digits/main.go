package Q58

import "strconv"

func addDigits(num int) int {
	for num >= 10 {
		tmp := 0
		for _, v := range strconv.Itoa(num) {
			tmp += int(v - 48)
		}
		num = tmp
	}
	return num
}

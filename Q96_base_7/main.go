package Q96

import "strconv"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	res, tmp := "", num

	if num < 0 {
		tmp *= -1
	}

	for tmp != 0 {
		res = strconv.Itoa(tmp%7) + res
		tmp /= 7
	}
	if num < 0 {
		return "-" + res
	}
	return res
}

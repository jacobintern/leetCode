package Q83

import (
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	if num1 == "0" {
		return num2
	} else if num2 == "0" {
		return num1
	}

	if l1, l2 := len(num1), len(num2); l1 > l2 {
		for i := 0; i < l1-l2; i++ {
			num2 = "0" + num2
		}
	} else {
		for i := 0; i < l2-l1; i++ {
			num1 = "0" + num1
		}
	}
	res := ""
	isCarry := false
	for i := len(num1) - 1; i >= 0; i-- {
		tmp := int(num1[i]+num2[i]) - 96
		if isCarry {
			tmp += 1
			isCarry = false
		}
		if tmp > 9 {
			isCarry = true
			tmp -= 10
		}
		res = strconv.Itoa(tmp) + res
	}

	if isCarry {
		res = "1" + res
	}

	return res
}

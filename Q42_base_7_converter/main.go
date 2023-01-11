package Q42

import "strconv"

func base7Converter(num int) string {
	tmpResult := 0
	tmpNum := 0
	if num < 10 {
		tmpNum = -num
	} else {
		tmpNum = num
	}

	for {
		if tmpNum < 7 {
			tmpResult = tmpResult + tmpNum
			break
		}
		tmpNum = tmpNum - 7
		tmpResult = tmpResult + 10
	}

	if num < 0 {
		return "-" + strconv.Itoa(tmpResult)
	} else {
		return strconv.Itoa(tmpResult)
	}
}

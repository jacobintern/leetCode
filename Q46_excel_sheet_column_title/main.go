package Q47

var list = map[int]string{
	1:  "A",
	2:  "B",
	3:  "C",
	4:  "D",
	5:  "E",
	6:  "F",
	7:  "G",
	8:  "H",
	9:  "I",
	10: "J",
	11: "K",
	12: "L",
	13: "M",
	14: "N",
	15: "O",
	16: "P",
	17: "Q",
	18: "R",
	19: "S",
	20: "T",
	21: "U",
	22: "V",
	23: "W",
	24: "X",
	25: "Y",
	26: "Z",
}

func convertToTitle(columnNumber int) string {
	res := ""
	tmp := columnNumber
	// 代表位置
	for tmp > 26 {
		v := tmp / 26

		if v == 0 {
			break
		}

		// 找出有幾個 26 就是那一個位數
		if v > 26 {
			res += list[v/26]
		} else {
			res += list[v]
		}

		tmp = ((v / 26) * 26) + v%26
	}

	b := columnNumber % 26
	// 最後加上餘數
	res += list[b]
	return res
}

func convertToTitle2(columnNumber int) string {
	rs := []byte{}
	for columnNumber > 0 {
		columnNumber--
		rs = append([]byte{byte(columnNumber%26 + 65)}, rs...)
		columnNumber /= 26
	}
	return string(rs)
}

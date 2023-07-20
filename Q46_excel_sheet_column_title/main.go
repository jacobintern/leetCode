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
	// 循環 1 + 26 直到不能算為止
	// 原理：1 + 26 + 26^2 + 26^3 + 26^4 + .....，以此類推
	// 當處理完 1 + 26 的時侯，取下一位數出來 26 + 26^2 然後除以 26 變成 1 + 26 再開始下一個循環的運算
	for columnNumber > 0 {
		// 減 1 後取餘數後取得目前位數的值
		columnNumber--
		res += list[columnNumber%26+1]
		// 除以 26 開始下一個循環
		columnNumber /= 26
	}
	return reverse(res)
}

func reverse(str string) string {
	var res string
	for i := len(str); i > 0; i-- {
		res += string(str[i-1])
	}
	return res
}

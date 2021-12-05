package Q39

var numList = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var keyList = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func intToRoman(num int) string {
	var str string
	for _, s := range keyList {

		if tmp := num / s; tmp > 0 {
			for j := 0; j < tmp; j++ {
				str += numList[s]
			}
			num -= s * tmp
		}

	}

	return str
}

package Q197

func shortestToChar(s string, c byte) []int {
	indexs, j := allIndex(s, c), 0
	res := make([]int, len(s))

	for i, v := range s {
		if c == byte(v) {
			res[i] = 0
		}
		if i > indexs[j] {
			j++
			res[i] = indexs[j] - i
		} else {
			res[i] = indexs[j] - i
		}
	}

	return res
}

func allIndex(s string, c byte) []int {
	var result []int
	for i := range len(s) {
		if s[i] == c {
			result = append(result, i)
		}
	}
	return result
}

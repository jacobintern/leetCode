package Q197

func shortestToChar(s string, c byte) []int {
	indexs := allIndex(s, c)
	n, l := len(s), len(indexs)-1
	res := make([]int, n)

	// 左
	j := 0
	for i := range s {
		if i > indexs[j] && j < l {
			j++
		}

		res[i] = absInt(indexs[j] - i)
	}

	// 右
	j = l
	for i := n - 1; i >= 0; i-- {
		if i < indexs[j] && j > 0 {
			j--
		}

		if tmp := absInt(i - indexs[j]); tmp < res[i] {
			res[i] = tmp
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

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

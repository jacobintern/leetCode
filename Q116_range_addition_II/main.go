package Q116

func maxCount(m, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}

	a, b := m, n
	for _, arr := range ops {
		if arr[0] < a {
			a = arr[0]
		}

		if arr[1] < b {
			b = arr[1]
		}
	}

	return a * b
}

package Q205

func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	result := make([][]int, n)

	for i := range result {
		result[i] = make([]int, m)
	}

	for i := range m {
		for j := range n {
			result[j][i] = matrix[i][j] // 轉置
		}
	}

	return result
}

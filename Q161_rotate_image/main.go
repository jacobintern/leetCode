package Q161

import "fmt"

func rotate(matrix [][]int) {
	l := len(matrix)
	// 對角兌換
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 左右兌換
	for i := 0; i < l; i++ {
		for j := 0; j < l/2; j++ {
			matrix[i][j], matrix[i][l-j-1] = matrix[i][l-j-1], matrix[i][j]
			fmt.Printf("[%d, %d], col: %d.", i, j, matrix[i][j])
		}
		fmt.Println()
	}
}

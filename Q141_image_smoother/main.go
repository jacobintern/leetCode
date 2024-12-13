package Q141

import "fmt"

func imageSmoother(img [][]int) [][]int {
	m, n := len(img)-1, len(img[0])-1

	newMatrix := make([][]int, m+1)

	for i, row := range img {
		newMatrix[i] = make([]int, n+1)
		for j, column := range row {
			fmt.Printf("[%d,%d] is %d", i, j, column)
			fmt.Println()
			avg := 0
			// 計算周圍平均
			if i == 0 && j == 0 { // 左上角
				avg = (img[i][j] + img[i][j+1] + img[i+1][j] + img[i+1][j+1]) / 4
			} else if i == 0 && j > 0 && j != n { // 上邊
				avg = (img[i][j-1] + img[i][j] + img[i][j+1] + img[i+1][j-1] + img[i+1][j] + img[i+1][j+1]) / 6
			} else if i == 0 && j == n { // 右上角
				avg = (img[i][j-1] + img[i][j] + img[i+1][j-1] + img[i+1][j]) / 4
			} else if i > 0 && j == 0 && i != m { // 左邊
				avg = (img[i-1][j] + img[i-1][j+1] + img[i][j] + img[i][j+1] + img[i+1][j] + img[i+1][j+1]) / 6
			} else if i > 0 && j == n && i != m { // 右邊
				avg = (img[i-1][j-1] + img[i-1][j] + img[i][j-1] + img[i][j] + img[i+1][j-1] + img[i+1][j]) / 6
			} else if i == m && j == 0 { // 左下角
				avg = (img[i-1][j] + img[i-1][j+1] + img[i][j] + img[i][j+1]) / 4
			} else if i == m && j > 0 && j != n { // 下邊
				avg = (img[i-1][j-1] + img[i-1][j] + img[i-1][j+1] + img[i][j-1] + img[i][j] + img[i][j+1]) / 6
			} else if i == m && j == n { // 右下角
				avg = (img[i-1][j-1] + img[i-1][j] + img[i][j-1] + img[i][j]) / 4
			} else { // 中間
				avg = (img[i-1][j-1] + img[i-1][j] + img[i-1][j+1] +
					img[i][j-1] + img[i][j] + img[i][j+1] +
					img[i+1][j-1] + img[i+1][j] + img[i+1][j+1]) / 9
			}
			fmt.Println(avg)
			newMatrix[i][j] = avg
		}
	}

	return newMatrix
}

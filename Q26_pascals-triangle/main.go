package Q26

func generate(numRows int) [][]int {
	arr := make([][]int, numRows)
	i := 0
	for i < numRows {
		arr2 := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				arr2[j] = 1
			} else {
				arr2[j] = arr[i-1][j-1] + arr[i-1][j]
			}
			arr[i] = arr2
		}
		i++
	}
	return arr
}

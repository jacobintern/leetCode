package Q27

func getRow(rowIndex int) []int {
    arr := make([]int, rowIndex)
	i := 0
	for i <= rowIndex {
		arr2 := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				arr2[j] = 1
			} else {
				arr2[j] = arr[j-1] + arr[j]
			}
		}
		arr = arr2
		i++
	}
	return arr
}

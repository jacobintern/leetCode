package Q200

func flipAndInvertImage(image [][]int) [][]int {
	r := reverse(image)
	i := invert(r)
	return i
}

func reverse(image [][]int) [][]int {
	for _, row := range image {
		left, right := 0, len(row)-1
		for left < right {
			row[left], row[right] = row[right], row[left] // 交換
			left++
			right--
		}
	}

	return image
}

func invert(image [][]int) [][]int {
	for _, row := range image {
		for j, column := range row {
			row[j] = column ^ 1
		}
	}

	return image
}

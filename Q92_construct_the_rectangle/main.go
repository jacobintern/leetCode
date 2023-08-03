package Q92

func constructRectangle(area int) []int {
	res, i := []int{}, 2
	m := make(map[int]int)

	for area > 0 {
		if area%2 == 0 {
			m[i]++
			area /= i
		}
	}
	return res
}

package Q199

func largeGroupPositions(s string) [][]int {
	s += " "
	res := [][]int{}
	start := 0

	for i, v := range s {
		if i == len(s) || byte(v) != s[start] {
			if i-start >= 3 {
				res = append(res, []int{start, i - 1})
			}
			start = i
		}
	}

	return res
}

package Q154

func permute(nums []int) [][]int {
	var res [][]int
	save([]int{}, nums, &res)

	return res
}

func save(cur, left []int, res *[][]int) {
	if len(left) == 0 {
		*res = append(*res, cur)
		return
	}

	for i, v := range left {
		cur := append(cur, v)
		var next []int
		// exclude me
		next = append(append(next, left[:i]...), left[i+1:]...)
		// // right
		// next = append(next, left[i+1:]...)

		save(cur, next, res)
	}
}

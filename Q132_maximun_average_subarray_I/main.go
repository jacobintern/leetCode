package Q132

func findMaxAverage(nums []int, k int) float64 {
	l, max := -1, 0
	for _, v := range nums[:k] {
		max += v
	}
	curM := max
	for l+k < len(nums)-1 {
		l++
		curM = curM - nums[l] + nums[l+k]
		if curM > max {
			max = curM
		}
	}

	return float64(max) / float64(k)
}

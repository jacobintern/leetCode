package Q91

func findMaxConsecutiveOnes(nums []int) int {
	max, cur := 0, 0

	for _, v := range nums {
		if v == 1 {
			cur += 1
		}
		if v == 0 {
			if cur > max {
				max = cur
			}
			cur = 0
		}
	}
	if cur > max {
		return cur
	}
	return max
}

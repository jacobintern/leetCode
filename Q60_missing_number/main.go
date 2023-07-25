package Q60

func missingNumber(nums []int) int {
	area := (len(nums) + 1) * len(nums) / 2

	for _, v := range nums {
		area -= v
	}

	return area
}

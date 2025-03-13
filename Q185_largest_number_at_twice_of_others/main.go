package Q185

import "slices"

func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return -1
	}
	var max, second int
	max = nums[0]
	for _, num := range nums[1:] {
		if num > second {
			second = num
		}
		if num > max {
			max, second = second, max
		}
	}

	if max >= second*2 {
		return slices.Index(nums, max)
	}

	return -1
}

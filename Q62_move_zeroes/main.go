package Q62

func moveZeroes(nums []int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[j] == 0 {
			// remove
			nums = append(nums[:j], nums[j+1:]...)
			// add
			nums = append(nums, 0)
		} else {
			j++
		}
	}

	return nums
}

// 記憶體位置變動了不被接受
func moveZeroes2(nums []int) []int {
	k := len(nums)
	zeroCount := 0
	for _, v := range nums {
		if v == 0 {
			zeroCount++
		} else {
			nums = append(nums, v)
		}
	}
	for zeroCount > 0 {
		nums = append(nums, 0)
		zeroCount--
	}
	nums = nums[k:]
	return nums
}

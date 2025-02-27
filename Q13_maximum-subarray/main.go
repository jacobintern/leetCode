package Q13

func maxSubArray(nums []int) int {
	// 初始值
	curr_max := nums[0]
	for i, curr := 1, nums[0]; i < len(nums); i++ {
		// 計算加總
		curr += nums[i]
		// 負數 or 小於前一數
		if curr < 0 || nums[i] > curr {
			curr = nums[i]
		}
		// 總數若小於前一數
		if curr_max < curr {
			curr_max = curr
		}
	}
	return curr_max
}

func maxSubArray2(nums []int) int {
	max, sum := nums[0], nums[0]
	for _, v := range nums[1:] {
		if sum < 0 {
			sum = v
		} else {
			sum += v
		}
		if max < sum {
			max = sum
		}
	}
	return max
}

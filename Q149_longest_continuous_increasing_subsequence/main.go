package Q149

func findLengthOfLCIS(nums []int) int {
	maxC, curC, curNum := 0, 1, nums[0]

	for _, num := range nums {
		if num > curNum {
			curC++
		} else {
			curC = 1
		}

		if curC > maxC {
			maxC = curC
		}

		curNum = num
	}

	return maxC
}

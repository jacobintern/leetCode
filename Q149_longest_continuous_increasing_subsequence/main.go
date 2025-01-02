package Q149

func findLengthOfLCIS(nums []int) int {
	maxC, curC, curNum := 0, 0, 0

	for _, num := range nums {
		if num > curNum {
			curC++
			if curC > maxC {
				maxC = curC
			}
		} else {
			curC = 1
		}
		curNum = num
	}

	return maxC
}

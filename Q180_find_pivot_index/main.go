package Q180

func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	leftSum := 0
	for i, num := range nums {
		rightSum := sum - num - leftSum
		if rightSum == leftSum {
			return i
		}
		leftSum += num
	}

	return -1
}

package Q137

func findErrorNums(nums []int) []int {
	for i := 0; i <= len(nums); i++ {
		if i+1 != nums[i] {
			return []int{nums[i], i + 1}
		}
	}

	return nil
}

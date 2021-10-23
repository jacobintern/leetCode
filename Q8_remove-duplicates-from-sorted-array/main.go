package Q8

func removeDuplicates(nums []int) int {
	var newArr []int
	for _, v := range nums {
		if !contains(newArr, v) {
			newArr = append(newArr, v)
		}
	}
	nums = newArr
	return len(nums)
}

func contains(nums []int, num int) bool {
	for _, v := range nums {
		if v == num {
			return true
		}
	}
	return false
}

func removeDuplicates2(nums []int) int {
	j := 0
	for i, k := 0, len(nums); i < k; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

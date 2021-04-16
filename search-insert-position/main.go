package Q11

func searchInsert(nums []int, target int) int {
	if v := IndexOf(nums, target); v >= 0 {
		return v
	}
	return InsertInt(nums, target)
}

func IndexOf(nums []int, num int) int {
	for i, v := range nums {
		if v == num {
			return i
		}
	}
	return -1
}

func InsertInt(nums []int, num int) int {
	r := 0
	if num == 0 || num < nums[0] {
		return r
	}
	if num > nums[len(nums)-1] {
		return len(nums)
	}
	for i, v := range nums {
		if v > num && num > nums[i-1] {
			r = i
		}
	}
	return r
}

package Q190

func minCapability(nums []int, k int) int {
	left, right := 1, int(1e9)
	for left < right {
		mid := (left + right) / 2
		if canRob(nums, k, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canRob(nums []int, k int, cap int) bool {
	count := 0
	i := 0
	for i < len(nums) {
		if nums[i] <= cap {
			count++
			i += 2 // 跳過相鄰的房屋
		} else {
			i++
		}
		if count >= k {
			return true
		}
	}
	return false
}

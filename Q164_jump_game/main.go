package Q164

func canJump(nums []int) bool {
	max := 0
	for i, v := range nums {
		if i > max {
			return false
		}
		if i+v > max {
			max = i + v
		}
	}
	return true
}

package Q159

func search(nums []int, target int) int {
	for p, v := range nums {
		if v == target {
			return p
		}
	}
	return -1
}

package Q9

func removeElement(nums []int, val int) int {
	r := 0
	for _, v := range nums {
		if v != val {
			nums[r] = v
			r++
		}
	}
	return r
}
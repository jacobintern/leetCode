package Q47

func majorityElement(nums []int) int {
	res := make(map[int]int)
	maxKey := 0
	maxValue := 0
	for _, v := range nums {
		res[v]++
		if res[v] > maxValue {
			maxKey = v
			maxValue = res[v]
		}
	}

	return maxKey
}

package Q160

func combinationSum4(nums []int, target int) int {
	res := make([]int, target+1)
	// default 0 ~ target -1
	for i := range res {
		res[i] = -1
	}

	return process(nums, target, res)
}

func process(nums []int, target int, res []int) int {
	if target == 0 {
		return 1
	}

	if res[target] != -1 {
		return res[target]
	}

	count := 0
	for _, num := range nums {
		if target >= num {
			count += process(nums, target-num, res)
		}
	}

	res[target] = count
	return count
}

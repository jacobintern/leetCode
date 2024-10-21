package Q137

func findErrorNums(nums []int) []int {
	dupArr := []int{}
	pre := 0
	for i, v := range nums {
		if i == 0 {
			pre = v
			continue
		}

		if v == pre {
			dupArr = append(dupArr, i, len(nums))
			break
		}

		pre = v
	}

	return dupArr
}

package Q31

func singleNumber(nums []int) int {
	
	elementMap := make(map[int]int)
	for _, num := range nums {
		elementMap[num]++
	}

	for i, ele := range elementMap {
		if ele == 1 {
			return i
		}
	}

	return 0
}

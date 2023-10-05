package Q123

func containsNearbyDuplicate(nums []int, k int) bool {
	numIndices := make(map[int]int)

	for i, num := range nums {
		if index, exists := numIndices[num]; exists && i-index <= k {
			return true
		}
		numIndices[num] = i
	}

	return false
}

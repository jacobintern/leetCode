package Q211

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sumAlice, sumBob := sum(aliceSizes), sum(bobSizes)
	delta := (sumBob - sumAlice) / 2

	bobSet := make(map[int]bool)
	for _, b := range bobSizes {
		bobSet[b] = true
	}

	for _, a := range aliceSizes {
		b := a + delta
		if bobSet[b] {
			return []int{a, b}
		}
	}

	return nil
}

func sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}

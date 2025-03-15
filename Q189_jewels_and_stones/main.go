package Q189

func numJewelsInStones(jewels string, stones string) int {
	// set map
	jewelsMap := make(map[rune]bool)
	for _, v := range jewels {
		jewelsMap[v] = true
	}

	// count
	count := 0
	for _, v := range stones {
		if jewelsMap[v] {
			count++
		}
	}

	return count
}

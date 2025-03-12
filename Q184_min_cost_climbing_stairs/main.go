package Q184

func minCostClimbingStairs(cost []int) int {
	first, second, n := cost[0], cost[1], len(cost)

	for i := 2; i < n; i++ {
		cur := cost[i] + min(first, second)
		first, second = second, cur
	}

	return min(first, second)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

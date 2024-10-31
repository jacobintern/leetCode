package Q140

func numTrees(n int) int {
	return traverse(1, n)
}

func traverse(left, right int) int {
	if left >= right {
		return 1
	}

	count := 0
	for i := left; i <= right; i++ {
		count += traverse(left, i-1) * traverse(i+1, right)
	}

	return count
}

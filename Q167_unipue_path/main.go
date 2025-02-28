package Q167

func uniquePaths(m int, n int) int {
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}
	return helper(0, 0, m, n, memo)
}

func helper(i, j, m, n int, memo [][]int) int {
	// 如果達到右下角， 1
	if i == m-1 && j == n-1 {
		return 1
	}

	// 如果越界， 0
	if i >= m || j >= n {
		return 0
	}

	// 如果已經計算過該位置，直接返回結果
	if memo[i][j] != 0 {
		return memo[i][j]
	}

	// 計算從右邊和下方兩個方向走的路徑數量
	memo[i][j] = helper(i+1, j, m, n, memo) + helper(i, j+1, m, n, memo)
	return memo[i][j]
}

package Q212

func surfaceArea(grid [][]int) int {
	n := len(grid)
	area := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				// 計算該位置的總表面積 (6 * 高度)
				area += 6 * grid[i][j]
				// 減去被遮擋的部分
				area -= (grid[i][j] - 1) * 2 // 自己遮擋自己
				if i > 0 {
					area -= min(grid[i][j], grid[i-1][j]) * 2 // 上下遮擋
				}
				if j > 0 {
					area -= min(grid[i][j], grid[i][j-1]) * 2 // 左右遮擋
				}
			}
		}
	}
	return area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package Q209

func projectionArea(grid [][]int) int {
    n := len(grid)
    topView, frontView, sideView := 0, 0, 0
    
    for i := range n {
        maxRow, maxCol := 0, 0
        for j := range n {
            if grid[i][j] > 0 {
                topView++ // 俯視圖
            }
            maxRow = max(maxRow, grid[i][j]) // 正視圖（行最大值）
            maxCol = max(maxCol, grid[j][i]) // 側視圖（列最大值）
        }
        frontView += maxRow
        sideView += maxCol
    }
    
    return topView + frontView + sideView
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

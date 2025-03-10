package Q182

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	if oldColor == newColor { // 避免無窮遞迴
		return image
	}

	dfs(image, sr, sc, oldColor, newColor)
	return image
}

func dfs(image [][]int, r, c, oldColor, newColor int) {
	if r < 0 || c < 0 || r >= len(image) || c >= len(image[0]) {
		return
	}

	if image[r][c] != oldColor {
		return
	}

	image[r][c] = newColor

	dfs(image, r+1, c, oldColor, newColor) // 下
	dfs(image, r-1, c, oldColor, newColor) // 上
	dfs(image, r, c+1, oldColor, newColor) // 右
	dfs(image, r, c-1, oldColor, newColor) // 左
}

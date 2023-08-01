package Q89

func islandPerimeter(grid [][]int) int {
	area := 0
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				area += 4

				if i == 0 && j == 0 {
					continue
				}

				if j > 0 {
					if grid[i][j-1] == 1 {
						area -= 2
					}
				}

				if i > 0 {
					if grid[i-1][j] == 1 {
						area -= 2
					}
				}
			}
		}
	}
	return area
}

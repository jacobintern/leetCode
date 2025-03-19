package Q195

func largestTriangleArea(points [][]int) float64 {
	maxX, minX, maxY, minY := 0, 50, 0, 50

	for _, point := range points {
		if tmpX := point[0]; tmpX > maxX {
			maxX = tmpX
		}
		if tmpY := point[1]; tmpY > maxY {
			maxY = tmpY
		}
		if tmpX := point[0]; tmpX < minX {
			minX = tmpX
		}
		if tmpY := point[1]; tmpY < minY {
			minY = tmpY
		}
	}

	area := float64((maxX - minX) * (maxY - minY))

	T1 := triangleArea(points[0][0], points[0][1], maxX, points[0][1], maxX, points[1][1])
	T2 := triangleArea(points[2][0], points[2][1], maxX, points[2][1], maxX, points[1][1])
	T3 := triangleArea(points[2][0], points[2][1], minY, points[2][1], minY, points[0][1])

	return area - T1 - T2 - T3
}

func triangleArea(x1, y1, x2, y2, x3, y3 int) float64 {
	base := float64(abs(x2 - x1))
	height := float64(abs(y3 - y1))
	return 0.5 * base * height
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

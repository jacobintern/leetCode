package Q37

func maxArea(height []int) int {
	max := 0
	i := 0
	j := len(height) - 1
	for i < j {
		if height[i] > height[j] {
			max = compareArea(height[j]*(j-i), max)
			j--
		} else {
			max = compareArea(height[i]*(j-i), max)
			i++
		}
	}

	return max
}

func compareArea(tmp, max int) int {
	if tmp > max {
		return tmp
	}
	return max
}

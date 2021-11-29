package Q37

// func maxArea(height []int) int {
// 	max := 0
// 	for i, h1 := range height {

// 		for j, h2 := range height {
// 			if i == j || i < j {
// 				continue
// 			}
// 			if h1 > h2 {
// 				max = compareArea(h2*(i-j), max)
// 				// fmt.Println("i is ", i, ", j is ", j, ", h1 is ", h1, ", h2 is", h2, "curArea is ", h2*(i-j))
// 			} else {
// 				max = compareArea(h1*(i-j), max)
// 				// fmt.Println("i is ", i, ", j is ", j, ", h1 is ", h1, ", h2 is", h2, "curArea is ", h1*(i-j))
// 			}
// 		}
// 	}

// 	return max
// }

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

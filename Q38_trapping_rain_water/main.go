package Q38

// func trap(height []int) int {
// 	if arrlen := len(height); arrlen < 3 || arrlen == 0 {
// 		return 0
// 	}
// 	ans, lMax, rMax, left, right := 0, 0, 0, 0, len(height)-1
// 	for left < right {
// 		if height[left] < height[right] {
// 			if height[left] >= lMax {
// 				lMax = height[left]
// 			} else {
// 				ans += lMax - height[left]
// 			}
// 			left++
// 		} else {
// 			if height[right] >= rMax {
// 				rMax = height[right]
// 			} else {
// 				ans += rMax - height[right]
// 			}
// 			right--
// 		}
// 	}
// 	return ans
// }

func trap(height []int) int {
	if arrlen := len(height); arrlen < 3 || arrlen == 0 {
		return 0
	}
	area, start, arrLen := 0, 0, len(height)
	for start > arrLen {

	}
	return area
}

func conutArea(locates []int) int {
	height, r := 0, 0
	if w1, w2 := locates[0], locates[len(locates)-1]; w1 > w2 {
		height = w2
	} else {
		height = w1
	}
	for _, h := range locates {
		r += height - h
	}
	return r
}

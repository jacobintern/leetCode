package Q92

import (
	"math"
)

func constructRectangle(area int) []int {
	l := int(math.Sqrt(float64(area)))
	for area%l != 0 {
		l++
	}

	if w := area / l; l < w {
		return []int{w, l}
	}

	return []int{l, area / l}
}

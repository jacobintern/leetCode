package Q78

import "fmt"

func readBinaryWatch(turnedOn int) []string {
	res := []string{}
	if turnedOn > 8 {
		return res
	}
	// 二進制 1 的數量等於亮的燈數
	list := make([]int, 60)

	for i := 1; i < 60; i++ {
		list[i] = list[i/2] + (i & 1)
	}

	// 取 H
	for h := 0; h < 12; h++ {
		m := 0
		// 取 M
		for m < 60 {
			if list[h]+list[m] == turnedOn {
				res = append(res, fmt.Sprintf("%d:%02d", h, m))
			}
			m++
		}
	}

	return res
}

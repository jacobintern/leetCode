package Q78

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch2(turnedOn int) []string {
	if turnedOn == 0 {
		return []string{"0:00"}
	}
	if turnedOn > 8 {
		return []string{}
	}

	res := []string{}
	// 先分配好數量再各自計算出結果，最後將結果組合起來
	// H 的次數必須小於 4
	// M 的次數必須小於 6
	m := make(map[int]int)
	for i := 3; i >= 0; i-- {
		if i > turnedOn || turnedOn-i >= 6 {
			continue
		}
		m[i] = turnedOn - i
	}

	return res
}

func readBinaryWatch(turnedOn int) []string {
	res := []string{}
	if turnedOn > 8 {
		return res
	}
	// 取 H
	for h := 0; h < 12; h++ {
		m := 0
		// 取 M
		for m < 60 {
			if bits.OnesCount(uint(h))+bits.OnesCount(uint(m)) == turnedOn {
				res = append(res, fmt.Sprintf("%d:%02d", h, m))
			}
			m++
		}
	}

	return res
}

package Q86

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	j := 0
	for _, v := range s {

		if j >= len(g) {
			break
		}

		if v >= g[j] {
			j++
		}
	}

	return j
}

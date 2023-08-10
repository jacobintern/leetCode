package Q97

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	tmp, m := make([]int, len(score)), make(map[int]int)
	copy(tmp, score)
	sort.Sort(sort.Reverse(sort.IntSlice(tmp)))

	for i, v := range tmp {
		m[v] = i + 1
	}

	res := []string{}

	for _, v := range score {
		switch m[v] {
		case 1:
			res = append(res, "Gold Medal")
		case 2:
			res = append(res, "Silver Medal")
		case 3:
			res = append(res, "Bronze Medal")
		default:
			res = append(res, strconv.Itoa(m[v]))
		}
	}

	return res
}

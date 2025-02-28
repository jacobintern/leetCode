package Q165

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}
	for _, v := range intervals[1:] {
		prev := merged[len(merged)-1]

		if v[0] <= prev[1] && prev[0] <= v[1] {
			prev[1] = max(prev[1], v[1])
		} else {
			merged = append(merged, v)
		}
	}

	return merged
}

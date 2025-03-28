package Q206

import "strconv"

func binaryGap(n int) int {
	binStr := strconv.FormatInt(int64(n), 2)
	maxGap, prev := 0, -1

	for i, ch := range binStr {
		if ch == '1' {
			if prev != -1 {
				maxGap = max(maxGap, i-prev)
			}
			prev = i
		}
	}

	return maxGap
}

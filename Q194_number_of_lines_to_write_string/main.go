package Q194

func numberOfLines(widths []int, s string) []int {
	lines, width := 1, 0

	for _, ch := range s {
		w := widths[ch-'a']
		if width+w > 100 {
			lines++
			width = w
		} else {
			width += w
		}
	}

	return []int{lines, width}
}

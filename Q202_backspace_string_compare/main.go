package Q202

func backspaceCompare(s string, t string) bool {
	ss, tt := "", ""
	ignore := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			ignore++
		} else if ignore != 0 {
			ignore--
		} else {
			ss = string(s[i]) + ss
		}
	}

	ignore = 0
	for i := len(t) - 1; i >= 0; i-- {
		if t[i] == '#' {
			ignore++
		} else if ignore != 0 {
			ignore--
		} else {
			tt = string(t[i]) + tt
		}
	}

	return ss == tt
}

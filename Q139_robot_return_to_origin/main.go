package Q139

func judgeCircle(moves string) bool {
	m := map[string]int{}
	m["h"] = 0
	m["v"] = 0
	for _, v := range moves {
		switch string(v) {
		case "U":
			m["v"] += 1
		case "D":
			m["v"] -= 1
		case "R":
			m["h"] += 1
		case "L":
			m["h"] -= 1
		}
	}
	if m["v"] == 0 && m["h"] == 0 {
		return true
	}

	return false
}

package Q104

func checkRecord(s string) bool {
	var L, A int
	for _, v := range s {
		if v == 'L' {
			L++
		} else {
			L = 0
		}
		if L == 3 {
			return false
		}
		if v == 'A' {
			A++
		}
		if A == 2 {
			return false
		}
	}
	return true
}

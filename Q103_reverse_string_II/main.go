package Q103

func reverseStr(s string, k int) string {
	sr := []byte(s)
	for l, r := 0, k-1; l < len(s); l, r = l+2*k, r+2*k {
		for i, j := l, min(len(sr)-1, r); i < j; i, j = i+1, j-1 {
			sr[i], sr[j] = sr[j], sr[i]
		}
	}
	return string(sr)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

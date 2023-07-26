package Q69

func reverseString(s []byte) []byte {
	for i, l := 0, len(s); i < l/2; i++ {
		left := s[i]
		right := s[l-i-1]

		if left != right {
			s[i] = right
			s[l-i-1] = left
		}
	}
	return s
}

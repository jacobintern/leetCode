package Q70

func reverseVowels(s string) string {
	left, right, res := 0, len(s)-1, []byte(s)

	for {
		if left == right || right < left {
			break
		} else if tmpl, tmpr := res[left], res[right]; contains(tmpl) && contains(tmpr) {
			res[left] = tmpr
			res[right] = tmpl
			left++
			right--
		} else if contains(tmpl) && !contains(tmpr) {
			right--
		} else if !contains(tmpl) && contains(tmpr) {
			left++
		} else if !contains(tmpl) && !contains(tmpr) {
			left++
			right--
		}
	}

	return string(res)
}

func contains(char byte) bool {
	arr := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, n := range arr {
		if char == n {
			return true
		}
	}
	return false
}

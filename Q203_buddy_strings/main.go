package Q203

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		seen := make(map[rune]bool)
		for _, ch := range s {
			if seen[ch] {
				return true // 只要有重複字母，就可以自己交換
			}
			seen[ch] = true
		}
		return false
	}

	// 找到不同的位置
	var first, second = -1, -1
	for i := range s {
		if s[i] != goal[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false // 超過兩個不同的字元
			}
		}
	}

	// 剛好兩個不同的位置，且它們可以交換
	return second != -1 && s[first] == goal[second] && s[second] == goal[first]
}

package Q101

func detectCapitalUse(word string) bool {
	count := 0

	for _, v := range word {
		if v < 97 {
			count++
		}
	}

	if count == len(word) || count == 0 || (count == 1 && word[0] < 97) {
		return true
	}

	return false
}

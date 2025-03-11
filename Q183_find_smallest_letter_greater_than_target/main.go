package Q183

func nextGreatestLetter(letters []byte, target byte) byte {
	var res, minGap byte
	minGap = 26
	for _, letter := range letters {
		if letter > target && letter-target < minGap {
			res = letter
			minGap = letter - target
		}
	}
	if minGap == 26 {
		return letters[0]
	}
	return res
}

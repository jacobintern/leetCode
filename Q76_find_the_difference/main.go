package Q76

func findTheDifference(s string, t string) byte {
	var sums, sumt rune

	for _, v := range s {
		sums += v
	}

	for _, v := range t {
		sumt += v
	}

	if sums > sumt {
		return byte(sums - sumt)
	}

	return byte(sumt - sums)
}

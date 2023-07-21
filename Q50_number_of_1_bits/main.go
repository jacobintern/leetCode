package Q50

func hammingWeight(num uint32) int {
	h := 0

	for num > 0 {
		if num%2 == 1 {
			h++
		}
		num /= 2
	}

	return h
}

func hammingWeight2(num uint32) int {
	h := 0
	for i := 0; i < 32; i++ {
		if num&1 == uint32(1) {
			h++
		}
		num = num >> 1
	}

	return h
}

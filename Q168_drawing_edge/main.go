package Q168

// 911
func drawingEdge(n int32) int64 {
	// get edge
	edge := int64(n * (n - 1) / 2)

	return powMod2(edge)
}

func powMod(exp int64) int64 {
	res, base := int64(1), int64(2)%mod

	for exp > 0 {
		if exp%2 == 1 {
			res = (res * base) % mod
		}
		base = (base * base) % mod
		exp /= 2
	}

	return res
}

func powMod2(exp int64) int64 {
	num := int64(1)

	for exp > 0 {
		num = (num * 2) % mod
		exp--
	}

	return num
}

const mod = 1000000007

package Q168

// 911
func drawingEdge(n int32) int32 {
	// get edge
	edge := int32((n * (n - 1)) / 2)
	newN := edge

	var res int32
	for newN >= 0 {
		res += combination(edge, newN)
		newN--
	}
	return res
}

func combination(n, r int32) int32 {
	if r > n {
		return 0
	}
	// 利用對稱性 C(n, r) = C(n, n-r)
	if r > n-r {
		r = n - r
	}
	result := int32(1)
	for i := int32(1); i <= r; i++ {
		result = result * (n - r + i) / i
		if result > mod {
			result = result % mod
		}
	}
	return result
}

const mod = 1000000007

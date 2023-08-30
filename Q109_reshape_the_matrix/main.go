package Q109

func matrixReshape(mat [][]int, r int, c int) [][]int {
	l, w := len(mat), len(mat[0])

	if l*w != r*c {
		return mat
	}

	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}

	for i := 0; i < l; i++ {
		for j := 0; j < w; j++ {
			l := i*w + j
			res[l/c][l%c] = mat[i][j]
		}
	}

	return res
}

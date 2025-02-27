package Q163

func spiralOrder(matrix [][]int) []int {
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	res := make([]int, 0, (r+1)*(b+1))
	for l <= r && t <= b {
		for i := l; i <= r; i++ { // 右
			res = append(res, matrix[t][i])
		}
		t++
		for i := t; i <= b; i++ { // 下
			res = append(res, matrix[i][r])
		}
		r--
		if t <= b {
			for i := r; i >= l; i-- { // 左
				res = append(res, matrix[b][i])
			}
			b--
		}
		if l <= r {
			for i := b; i >= t; i-- { // 上
				res = append(res, matrix[i][l])
			}
			l++
		}
	}
	return res
}

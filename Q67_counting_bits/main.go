package Q67

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[i/2] + 1
		}
	}
	return res
}

// func countBits(n int) []int {
// 	var res []int
// 	for i := 0; i <= n; i++ {
// 		sum := 0
// 		for _, v := range getBytesArr(i) {
// 			sum += v
// 		}
// 		res = append(res, sum)
// 	}
// 	return res
// }

// func getBytesArr(n int) []int {
// 	arr := []int{}
// 	for n > 0 {
// 		arr = append(arr, n%2)
// 		n /= 2
// 	}
// 	return arr
// }

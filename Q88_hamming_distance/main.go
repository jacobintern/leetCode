package Q88

func hammingDistance(x int, y int) int {
	arrx := decimalToBinary(x)
	arry := decimalToBinary(y)
	if a, b := len(arrx), len(arry); a > b {
		arry = amendZero(arry, a-b)
	} else {
		arrx = amendZero(arrx, b-a)
	}
	res := 0
	for i := 0; i < len(arrx); i++ {
		res += arrx[i] ^ arry[i]
	}

	return res
}

func decimalToBinary(num int) []int {
	res := []int{}
	for num > 0 {
		res = append([]int{num % 2}, res...)
		num /= 2
	}
	return res
}

func amendZero(nums []int, num int) []int {
	for num > 0 {
		nums = append([]int{0}, nums...)
		num--
	}
	return nums
}

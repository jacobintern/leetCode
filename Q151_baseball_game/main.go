package Q151

import "strconv"

func calPoints(operations []string) int {
	var (
		arr []int
		res int
	)

	for _, v := range operations {
		arr = process(v, arr)
	}

	for _, v := range arr {
		res += v
	}

	return res
}

func process(s string, arr []int) []int {
	l := len(arr)
	switch s {
	case "+":
		arr = append(arr, arr[l-1]+arr[l-2])
	case "D":
		arr = append(arr, arr[l-1]*2)
	case "C":
		arr = arr[:l-1]
	default:
		v, _ := strconv.Atoi(s)
		arr = append(arr, v)
	}
	return arr
}

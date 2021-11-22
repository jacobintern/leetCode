package Q32

import (
	"sort"
)

func findDisappearedNumbers(nums []int) []int {
	var result []int
	sort.Ints(nums)
	for i, numsLen := 1, len(nums); i <= numsLen; i++ {
		if !Contains(nums, i) {
			// fmt.Println("i is ", i, ", num is ", nums[i-1])
			result = append(result, i)
		}
	}
	return result
}

func Contains(arr []int, x int) bool {
	for _, n := range arr {
		if x == n {
			return true
		}
	}
	return false
}

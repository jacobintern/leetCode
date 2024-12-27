package Q146

import (
	"sort"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	l := len(nums)

	// 正數最大值
	max1 := nums[l-1] * nums[l-2] * nums[l-3]
	// 兩負數乘正數
	max2 := nums[0] * nums[1] * nums[l-1]

	// 回傳最大值
	if max1 > max2 {
		return max1
	}
	return max2
}

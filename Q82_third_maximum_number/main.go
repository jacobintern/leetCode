package Q82

import "sort"

func thirdMax(nums []int) int {
	sort.Ints(nums)
	res := []int{}
	tmp := 0
	res = append(res, nums[0])
	for _, v := range nums {
		if v != res[tmp] {
			res = append(res, v)
			tmp++
		}
	}

	if len(res) < 3 {
		return res[len(res)-1]
	}

	return res[len(res)-3]
}

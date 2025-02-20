package Q40

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	res, n := [][]int{}, len(nums)
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r, cur := i+1, n-1, 0-nums[i]
		for l < r {
			if nums[l]+nums[r] == cur {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if nums[l]+nums[r] > cur {
				r--
			} else {
				l++
			}
		}
	}
	return res
}

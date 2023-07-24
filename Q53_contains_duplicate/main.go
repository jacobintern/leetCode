package Q53

import (
	"sort"
)

func containsDuplicate3(nums []int) bool {
	if len(nums) == 1 {
		return false
	}

	sort.Ints(nums)
	pre := -1
	for _, v := range nums {
		if pre == v {
			return true
		}
		pre = v
	}
	return false
}

func containsDuplicate2(nums []int) bool {
	m := make(map[int]bool)

	for _, v := range nums {
		if m[v] {
			return true
		}
		m[v] = true
	}

	return false
}

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i, v := range nums {
		if i == 0 {
			continue
		} else if nums[i-1] == v {
			return true
		}
	}
	return false
}

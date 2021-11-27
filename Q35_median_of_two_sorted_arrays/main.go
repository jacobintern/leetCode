package Q35

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := mergeTwoArray(nums1, nums2)
	if arrLen := len(arr); arrLen%2 > 0 {
		return float64(arr[(arrLen / 2)])
	} else {
		return float64((arr[(arrLen/2)-1] + arr[(arrLen/2)])) / 2
	}
}

func mergeTwoArray(nums1 []int, nums2 []int) []int {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	return nums1
}

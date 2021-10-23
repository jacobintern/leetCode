package Q20

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	for total, m, n := m+n-1, m-1, n-1; total >= 0; total-- {
		if m < 0 || (n >= 0 && nums1[m] < nums2[n]) {
			nums1[total] = nums2[n]
			n--
		} else {
			nums1[total] = nums1[m]
			m--
		}
	}
	return nums1
}

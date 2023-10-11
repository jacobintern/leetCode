package Q127

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		m[nums1[i]]++
	}

	var res []int

	for i := 0; i < len(nums2); i++ {
		if count, ok := m[nums2[i]]; ok && count > 0 {
			res = append(res, nums2[i])
			m[nums2[i]]--
		}
	}

	return res
}

package Q71

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	res := []int{}
	for _, v := range nums1 {
		_, ok := m[v]
		if !ok {
			m[v]++
		}
	}

	for _, v := range nums2 {
		value, ok := m[v]
		if ok && value == 1 {
			m[v]++
		}
	}

	for i, v := range m {
		if v == 2 {
			res = append(res, i)
		}
	}
	return res
}

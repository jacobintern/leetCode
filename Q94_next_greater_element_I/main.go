package Q94

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var res []int

	for _, v := range nums1 {
		stamp, tmp := -1, 0
		for i := 0; i < len(nums2); i++ {
			if v == nums2[i] {
				stamp = i
			}

			if stamp > -1 && nums2[i] > v {
				tmp = nums2[i]
				break
			}
		}
		if stamp == -1 || tmp == 0 {
			res = append(res, -1)
		} else if stamp > -1 && tmp > 0 {
			res = append(res, tmp)
		}
	}
	return res
}

package Q117

func findRestaurant(list1 []string, list2 []string) []string {
	if len(list1) == 1 && len(list2) == 1 {
		return list1
	}
	i, j, min := 0, 0, len(list1)+len(list2)
	res := []string{}
	for i < len(list1) || j < len(list2) {

		if list1[i] == list2[j] {
			if i+j == min {
				res = append(res, list1[i])
			}
			if i+j < min {
				res = []string{list1[i]}
				min = i + j
			}
		}

		if j < len(list2) {
			j++
		}
		if i+1 < len(list1) && j == len(list2) {
			i++
			j = 0
		} else if i+1 == len(list1) && j == len(list2) {
			break
		}
	}

	return res
}

func findRestaurant2(list1 []string, list2 []string) []string {
	m := make(map[string]int)
	res := []string{}
	min := len(list1) + len(list2)

	for i := 0; i < len(list1); i++ {
		m[list1[i]] = i
	}

	for i := 0; i < len(list2); i++ {
		if j, ok := m[list2[i]]; ok && i+j < min {
			min = i + j
			res = []string{list2[i]}
		} else if ok && i+j == min {
			res = append(res, list2[i])
		}
	}
	return res
}

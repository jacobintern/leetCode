package Q172

func findShortestSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	// map
	m := make(map[int]int)
	max := 0
	for _, v := range nums {
		_, exists := m[v]
		if exists {
			m[v] += 1
		} else {
			m[v] = 1
		}

		if m[v] > max {
			max = m[v]
		}
	}

	// find degree
	degreeMap := make(map[int]int)
	for num, count := range m {
		if count == max {
			degreeMap[num] = count
		}
	}

	min := 50000
	for degree := range degreeMap {
		l, r := 0, len(nums)-1
		for l <= r {
			if nums[l] == degree && nums[r] == degree {
				if c := r - l + 1; c < min {
					min = c
				}
				break
			}
			if nums[l] != degree {
				l++
			}

			if nums[r] != degree {
				r--
			}
		}
	}

	return min
}

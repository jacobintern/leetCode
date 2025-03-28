package Q204

func lemonadeChange(bills []int) bool {
	cash := map[int]int{5: 0, 10: 0}

	for _, bill := range bills {
		if bill == 5 {
			cash[5]++
		} else if bill == 10 {
			if cash[5] == 0 {
				return false
			}
			cash[5]--
			cash[10]++
		} else {
			if cash[10] > 0 && cash[5] > 0 {
				cash[10]--
				cash[5]--
			} else if cash[5] >= 3 {
				cash[5] -= 3
			} else {
				return false
			}
		}
	}
	return true
}

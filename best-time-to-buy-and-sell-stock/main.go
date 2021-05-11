package Q28

func maxProfit(prices []int) int {
	r := 0
	for i, v := range prices {
		for j := i; j < len(prices); j++ {
			if i != j && j > i {
				tmp := prices[j] - v
				if r < tmp {
					r = tmp
				}
			}
		}
	}
	return r
}

func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxProfit := 0
	minPrice := prices[0]
	for i := range prices {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}

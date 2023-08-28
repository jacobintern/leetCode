package Q29

func maxProfit(prices []int) int {
	buy := 0
	sell := 0
	profit := 0
	n := len(prices)
	for buy < n && sell < n {
		for buy+1 < n && prices[buy+1] < prices[buy] {
			buy++
		}
		sell = buy
		for sell+1 < n && prices[sell+1] > prices[sell] {
			sell++
		}

		profit += prices[sell] - prices[buy]
		buy = sell + 1
	}
	return profit
}

// func main() {
// 	a := []int{7, 1, 5, 3, 6, 4}
// 	fmt.Println(maxProfit(a))
// }

// func maxProfit(prices []int) int {
// 	r := 0
// 	for i, v := range prices {
// 		for j := i; j < len(prices); j++ {
// 			if i != j && j > i {
// 				tmp := prices[j] - v
// 				if r < tmp {
// 					r = tmp
// 				}
// 			}
// 		}
// 	}
// 	return r
// }

// func maxProfit2(prices []int) int {
// 	if len(prices) == 0 {
// 		return 0
// 	}
// 	maxProfit := 0
// 	minPrice := prices[0]
// 	for i := range prices {
// 		if prices[i] < minPrice {
// 			minPrice = prices[i]
// 		} else if prices[i]-minPrice > maxProfit {
// 			maxProfit = prices[i] - minPrice
// 		}
// 	}
// 	return maxProfit
// }

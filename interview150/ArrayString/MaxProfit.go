package interview150

import "fmt"

// Best Time to Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/?envType=study-plan-v2&envId=top-interview-150
func MaxProfit() {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {
	min, max, profit := prices[0], prices[0], 0
	for _, v := range prices {
		if v < min {
			if profit < max-min {
				profit = max - min
			}
			min, max = v, v
		} else {
			if max < v {
				max = v
			}
		}
	}
	if profit < max-min {
		profit = max - min
	}
	return profit
}

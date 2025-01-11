package interview150

import "fmt"

// Best Time to Buy and Sell Stock II
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/?envType=study-plan-v2&envId=top-interview-150
func MaxProfit2() {
	fmt.Println(maxProfit2([]int{1, 2, 3, 4, 5}))
}

func maxProfit2(prices []int) int {
	min, max, profit := prices[0], prices[0], 0

	for _, v := range prices {
		if v < max {
			profit += max - min
			max, min = v, v
		} else {
			max = v
		}
	}
	profit += max - min
	return profit
}

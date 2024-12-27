package interview150

import (
	"fmt"
)

// Coin Change
// https://leetcode.com/problems/coin-change/submissions/1488721803/?envType=study-plan-v2&envId=top-interview-150
func CoinChange() {

	fmt.Println(coinChange([]int{2}, 1))
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	list := make([]int, amount+1)
	for _, v := range coins {
		if v < len(list) {
			list[v] = 1
		}
	}
	fmt.Println(list)

	for i := range list {
		for _, coin := range coins {
			if coin < i && list[i-coin] > 0 {
				if list[i] == 0 {
					list[i] = list[i-coin] + 1
				} else {
					if list[i] > list[i-coin]+1 {
						list[i] = list[i-coin] + 1
					}
				}
			}
		}
	}
	fmt.Println(list)
	if list[amount] == 0 {
		return -1
	}
	return list[amount]
}

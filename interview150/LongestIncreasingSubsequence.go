package interview150

import (
	"fmt"
)

func LengthOfLIS() {
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
}

func lengthOfLIS(nums []int) int {
	max := 1
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := i; j >= 0; j-- {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				if dp[i] > max {
					max = dp[i]
				}
			}
		}
	}
	return max
}

package main

import (
	"fmt"

	"rellik24.com/m/v2/interview150"
)

func MaxSubArray() {
	fmt.Println(maxSubArray([]int{-2, -1}))
}

func maxSubArray(nums []int) int {
	max, sum := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		sum = interview150.GetMax(sum+nums[i], nums[i])
		max = interview150.GetMax(max, sum)
	}

	return max
}

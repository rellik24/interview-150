package interview150

import (
	"fmt"
)

// Minimum Size Subarray Sum
// https://leetcode.com/problems/minimum-size-subarray-sum/description/?envType=study-plan-v2&envId=top-interview-150
func MinSubArrayLen() {
	fmt.Println(minSubArrayLen(71, []int{2, 3, 1, 2, 4, 3}))
}

func minSubArrayLen(target int, nums []int) int {
	min := len(nums) + 1

	left, sum := 0, 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			if min > right-left+1 {
				min = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	if min > len(nums) {
		return 0
	}
	return min
}

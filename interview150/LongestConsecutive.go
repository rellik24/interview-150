package interview150

import (
	"fmt"
	"sort"
)

// Longest Consecutive Sequence
// https://leetcode.com/problems/longest-consecutive-sequence/description/?envType=study-plan-v2&envId=top-interview-150
func LongestConsecutive() {
	fmt.Println(longestConsecutive([]int{0, 1, 1, 2}))
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)

	max := 1
	tmp := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == tmp+1 {
			count++
		} else if nums[i] == tmp {
			continue
		} else {
			if count > max {
				max = count
			}
			count = 1
		}
		tmp = nums[i]
	}
	if count > max {
		max = count
	}
	return max
}

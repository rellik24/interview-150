package interview150

import (
	"fmt"
)

// Single Number
// https://leetcode.com/problems/single-number/?envType=study-plan-v2&envId=top-interview-150
func SingleNumber() {
	fmt.Println(singleNumber([]int{1, 2, 1}))
}

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num // XOR 運算
	}
	return result
}

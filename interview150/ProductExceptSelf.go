package interview150

import (
	"fmt"
)

// Product of Array Except Self
// https://leetcode.com/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=top-interview-150
func ProductExceptSelf() {
	fmt.Println(productExceptSelf([]int{-1, 1, 0, 3, -3}))
}

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	countZ, idx := 0, 0
	max := 1
	for i, v := range nums {
		if v == 0 {
			countZ++
			idx = i
			continue
		} else {
			max = max * v
		}
	}
	if countZ > 1 {

	} else if countZ == 1 {
		result[idx] = max
	} else {
		for i, v := range nums {
			if v != 0 {
				result[i] = max / v
			}
		}
	}

	return result
}

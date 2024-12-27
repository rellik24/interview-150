package interview150

import (
	"fmt"
)

// Majority Element
// https://leetcode.com/problems/majority-element/?envType=study-plan-v2&envId=top-interview-150
func MajorityElement() {
	fmt.Println(majorityElement([]int{1}))
}

func majorityElement(nums []int) int {
	max, count := 0, 0
	m := make(map[int]int)
	for _, v := range nums {
		if val, ok := m[v]; !ok {
			m[v] = 1
		} else {
			val += 1
			m[v] = val
		}
		if m[v] > count {
			count = m[v]
			max = v
		}
	}
	return max
}

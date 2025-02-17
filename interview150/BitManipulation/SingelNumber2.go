package interview150

import "fmt"

// Single Number II
// https://leetcode.com/problems/single-number-ii/description/?envType=study-plan-v2&envId=top-interview-150
func SingleNumber2() {
	fmt.Println(singleNumber2([]int{2, 2, 3, 2}))
}

func singleNumber2(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

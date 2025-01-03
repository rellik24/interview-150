package interview150

import "fmt"

// Contains Duplicate II
// https://leetcode.com/problems/contains-duplicate-ii/description/?envType=study-plan-v2&envId=top-interview-150
func ContainsNearbyDuplicate() {
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	for i, v := range nums {
		if idx, exist := m[v]; exist {
			if i-idx <= k {
				return true
			}
		}
		m[v] = i
	}
	return false
}

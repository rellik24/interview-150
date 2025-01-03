package interview150

import "fmt"

// Two Sum II - Input Array Is Sorted
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/?envType=study-plan-v2&envId=top-interview-150
func TwoSum() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		if idx, ok := m[target-v]; ok && i != idx {
			return []int{idx, i}
		}
		m[v] = i
	}

	return []int{}
}

package interview150

import "fmt"

// Two Sum II - Input Array Is Sorted
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/?envType=study-plan-v2&envId=top-interview-150
func TwoSum() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for idx, v := range numbers {
		m[v] = idx
	}
	for idx, v := range numbers {
		if val, ok := m[target-v]; ok {
			return []int{idx + 1, val + 1}
		}
	}
	return []int{}
}

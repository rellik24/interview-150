package interview150

import "fmt"

// Remove Duplicates from Sorted Array II
// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/?envType=study-plan-v2&envId=top-interview-150
func RemoveDuplicates2() {
	fmt.Println(removeDuplicates2([]int{1, 1, 1, 2, 2, 3}))
}

func removeDuplicates2(nums []int) int {
	l := len(nums) - 1
	i := 0

	m := make(map[int]int)
	for i <= l {

		val := nums[i]
		if count, ok := m[val]; ok {
			if count < 2 {
				m[val] = count + 1
			} else {
				nums = append(nums[:i], nums[i+1:]...)
				l--
				continue
			}
		} else {
			m[val] = 1
		}
		i++
	}
	return len(nums)
}

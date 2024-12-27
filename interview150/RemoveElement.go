package interview150

import "fmt"

// RemoveElement
// https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150
func RemoveElement() {
	fmt.Println(removeElement([]int{3, 3, 2, 2, 3}, 3))
}

func removeElement(nums []int, val int) int {
	l := len(nums) - 1
	i := 0
	for i <= l {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			l--
			continue
		}
		i++
	}
	return len(nums)
}

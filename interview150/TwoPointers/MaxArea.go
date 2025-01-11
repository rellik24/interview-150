package interview150

import "fmt"

// Container With Most Water
// https://leetcode.com/problems/container-with-most-water/description/?envType=study-plan-v2&envId=top-interview-150
func MaxArea() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	l := len(height) - 1
	left, right := 0, l // index
	max := 0

	for l >= 1 {
		if height[left] < height[right] {
			if max < height[left]*l {
				max = height[left] * l
			}
			left++
		} else {
			if max < height[right]*l {
				max = height[right] * l
			}
			right--
		}
		l--
	}
	return max
}

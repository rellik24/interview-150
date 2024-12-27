package interview150

import (
	"fmt"
	"sort"
)

// 3Sum
// https://leetcode.com/problems/3sum/description/?envType=study-plan-v2&envId=top-interview-150
func ThreeSum() {
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
}

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	l := len(nums) - 1
	if len(nums) < 3 || nums[0] > 0 || nums[len(nums)-1] < 0 {
		return [][]int{}
	}
	for i := 0; i < l; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := nums[i] * -1
		// fmt.Println(nums[i])
		left := i + 1
		right := l
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// move forward if left is duplicate
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// move forward if right is duplicate
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < target {
				left++
			} else if sum > target {
				right--
			}
		}
	}
	return result
}

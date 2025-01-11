package interview150

import (
	"fmt"
	"strconv"
)

// Summary Ranges
// https://leetcode.com/problems/summary-ranges/description/?envType=study-plan-v2&envId=top-interview-150
func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	var result []string
	start, last := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == last+1 {
			last = nums[i]
		} else {
			if start == last {
				result = append(result, strconv.Itoa(start))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", start, last))
			}
			start, last = nums[i], nums[i]
		}
	}
	if start == last {
		result = append(result, strconv.Itoa(start))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", start, last))
	}
	return result
}

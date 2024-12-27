package interview150

import "fmt"

func Rob() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

func rob(nums []int) int {
	var max int
	list := make([]int, len(nums))
	for i, v := range nums {
		if i == 0 || i == 1 {
			list[i] = v
		} else if i == 2 {
			list[i] = list[0] + nums[2]
		} else {
			tmp := list[i-2]
			if tmp < list[i-3] {
				tmp = list[i-3]
			}
			list[i] = tmp + nums[i]
		}
		if list[i] > max {
			max = list[i]
		}

	}
	return max
}

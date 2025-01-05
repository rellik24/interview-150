package interview150

import (
	"fmt"
	"sort"
)

// Minimum Number of Arrows to Burst Balloons
// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/description/?envType=study-plan-v2&envId=top-interview-150
// 用 end 排序，每次選擇最早結束的位置作為射箭點，一旦出現新的 start 超過 end 表示要新射箭 (aka 更新)
func FindMinArrowShots() {
	fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
}

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	count := 1
	end := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] <= end {
			continue
		} else {
			count++
			end = points[i][1]
		}
	}

	return count
}

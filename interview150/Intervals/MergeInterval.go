package interview150

import (
	"fmt"
	"sort"

	"rellik24.com/m/v2/interview150"
)

// Merge Intervals
// https://leetcode.com/problems/merge-intervals/description/?envType=study-plan-v2&envId=top-interview-150
func Merge() {
	fmt.Println(mergeIntervals([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	var result [][]int
	sort.Slice(intervals, func(i, j int) bool { // 排序
		return intervals[i][0] < intervals[j][0]
	})
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]

		idx, val := len(result)-1, result[len(result)-1] // 只跟最後的比較好
		if interval[0] > val[1] {
			result = append(result, interval)
		} else {
			result[idx] = []int{interview150.GetMin(val[0], interval[0]), interview150.GetMax(val[1], interval[1])}
		}
	}
	return result
}

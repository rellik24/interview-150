package interview150

import (
	"fmt"
)

// Insert Interval
// https://leetcode.com/problems/insert-interval/description/?envType=study-plan-v2&envId=top-interview-150
// NOTE: intervals 本身不重疊，簡單串接就好，只要把 newInterval 疊加好就行
func InsertInterval() {
	fmt.Println(insertInterval([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
}

func insertInterval(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	i := 0

	// not overlap with newInterval
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// overlap with newInterval
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = getMin(intervals[i][0], newInterval[0])
		newInterval[1] = getMax(intervals[i][1], newInterval[1])
		i++
	}
	result = append(result, newInterval)

	// not overlap with newInterval
	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}
	return result
}

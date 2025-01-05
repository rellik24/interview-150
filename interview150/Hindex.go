package interview150

import (
	"fmt"
	"sort"
)

// H-Index
// https://leetcode.com/problems/h-index/description/?envType=study-plan-v2&envId=top-interview-150
func HIndex() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
}

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] >= i+1 {
			return i + 1
		}
	}
	return 0
}

package interview150

import (
	"sort"
)

// MergeSortedArray
// https://leetcode.com/problems/merge-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150
func MergeSortedArray() {
	var nums1 = []int{1, 2, 3, 0, 0, 0}
	var nums2 = []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	nums1 = nums1[:m]
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
}

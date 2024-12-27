package interview150

// Remove Duplicates from Sorted Array
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&envId=top-interview-150
func RemoveDuplicates() {
	removeDuplicates([]int{1, 1, 1, 2, 2, 3})
}

func removeDuplicates(nums []int) int {
	l := len(nums) - 1
	i := 0

	m := make(map[int]int)
	for i <= l {

		val := nums[i]
		if _, ok := m[val]; ok {
			nums = append(nums[:i], nums[i+1:]...)
			l--
			continue
		} else {
			m[val] = 1
		}
		i++
	}
	return len(nums)
}

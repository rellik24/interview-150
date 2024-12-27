package interview150

// Rotate Array
// https://leetcode.com/problems/rotate-array/description/?envType=study-plan-v2&envId=top-interview-150
func Rotate() {

	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	t := l - k
	copy(nums, append(nums[t:], nums[:t]...))
}

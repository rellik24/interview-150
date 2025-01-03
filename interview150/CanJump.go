package interview150

import "fmt"

// Jump Game
// https://leetcode.com/problems/jump-game/description/?envType=study-plan-v2&envId=top-interview-150
func CanJump() {
	fmt.Println(canJump([]int{0, 2, 3, 1, 1, 4}))
}

func canJump(nums []int) bool {
	l := len(nums)
	list := make([]bool, l)
	list[0] = true
	for i, v := range nums {
		for j := 1; j <= v; j++ {
			if i+j <= l-1 && list[i] {
				list[i+j] = true
			} else {
				break
			}
		}
	}
	return list[l-1]
}

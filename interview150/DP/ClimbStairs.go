package interview150

import (
	"fmt"
)

// Climbing Stairs
// https://leetcode.com/problems/climbing-stairs/description/?envType=study-plan-v2&envId=top-interview-150
func ClimbStairs() {
	fmt.Println(climbStairs(4))
}

func climbStairs(n int) int {
	stairs := []int{1, 2}
	for i := 2; i < n; i++ {
		stairs = append(stairs, 0)
	}
	fmt.Println(stairs)

	for i := 2; i < n; i++ {
		stairs[i] = stairs[i-1] + stairs[i-2]
	}
	fmt.Println(stairs)
	return stairs[n-1]
}

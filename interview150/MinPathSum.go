package interview150

import (
	"fmt"
)

// Minimum Path Sum
// https://leetcode.com/problems/minimum-path-sum/description/?envType=study-plan-v2&envId=top-interview-150
func MinPathSum() {
	var grid [][]int
	grid = append(grid, []int{1, 4, 8, 6, 2, 2, 1, 7})
	grid = append(grid, []int{4, 7, 3, 1, 4, 5, 5, 1})
	grid = append(grid, []int{8, 8, 2, 1, 1, 8, 0, 1})
	grid = append(grid, []int{8, 9, 2, 9, 8, 0, 8, 9})
	grid = append(grid, []int{5, 7, 5, 7, 1, 8, 5, 5})
	grid = append(grid, []int{7, 0, 9, 4, 5, 6, 5, 6})
	grid = append(grid, []int{4, 9, 9, 7, 9, 1, 9, 0})
	fmt.Println(minPathSum(grid))
}

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range grid {
		dp[i] = make([]int, len(grid[i]))
	}

	dp[0][0] = grid[0][0]
	for i, v := range grid {
		for j := range v {
			if i == 0 && j > 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if i > 0 && j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else if i > 0 && j > 0 {
				if dp[i][j-1] < dp[i-1][j] {
					dp[i][j] = grid[i][j] + dp[i][j-1]
				} else {
					dp[i][j] = grid[i][j] + dp[i-1][j]
				}
			}
		}
	}
	fmt.Println(dp)

	return dp[len(grid)-1][len(grid[0])-1]
}

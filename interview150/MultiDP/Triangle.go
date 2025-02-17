package interview150

import (
	"fmt"

	"rellik24.com/m/v2/interview150"
)

func MinimumTotal() {
	var board [][]int
	board = append(board, []int{2})
	board = append(board, []int{3, 4})
	board = append(board, []int{6, 5, 7})
	board = append(board, []int{4, 1, 8, 3})
	fmt.Println(minimumTotal(board))
}

func minimumTotal(triangle [][]int) int {
	var dp [][]int
	for _, v := range triangle {
		tmp := make([]int, len(v))
		dp = append(dp, tmp)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := range triangle[i] {
			if j == 0 {
				dp[i][j] = triangle[i][j] + dp[i-1][0]
			} else if j == len(triangle[i])-1 {
				dp[i][j] = triangle[i][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = interview150.GetMin(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}
		}
	}
	minPath := dp[len(dp)-1][0]
	for _, v := range dp[len(dp)-1] {
		if v < minPath {
			minPath = v
		}
	}
	return minPath
}

package interview150

import "fmt"

// Maximal Square
// https://leetcode.com/problems/maximal-square/description/?envType=study-plan-v2&envId=top-interview-150
// NOTE: 除了 Index: j==0 or i==0，每個格子都作為正方形的最右下角去疊加，並將左、上、左上選最小+1（如果最小是0表示非正方形）
// 最大的值就是 “邊長”
func MaximalSquare() {
	var matrix [][]byte
	matrix = append(matrix, []byte{'1', '0', '1', '0', '0'})
	matrix = append(matrix, []byte{'1', '0', '1', '1', '1'})
	matrix = append(matrix, []byte{'1', '1', '1', '1', '1'})
	matrix = append(matrix, []byte{'1', '0', '0', '1', '0'})

	fmt.Println(maximalSquare(matrix))
}

func maximalSquare(matrix [][]byte) int {
	var dp [][]int
	for _, v := range matrix {
		var tmp []int
		for _, val := range v {
			tmp = append(tmp, int(val-'0'))
		}
		dp = append(dp, tmp)
	}
	// fmt.Println(dp)
	for i, val := range dp {
		for j := range val {
			if i == 0 || j == 0 {

			} else if dp[i][j] == 1 {
				dp[i][j] = getMin(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	fmt.Println(dp)
	max := 0
	for _, val := range dp {
		for _, v := range val {
			if v > max {
				max = v
			}
		}
	}

	return max * max
}

func getMin(x, y, z int) int {
	min := x
	if y < min {
		min = y
	}
	if z < min {
		min = z
	}
	return min
}

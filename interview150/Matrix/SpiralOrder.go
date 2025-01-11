package interview150

import "fmt"

// Spiral Matrix
// https://leetcode.com/problems/spiral-matrix/?envType=study-plan-v2&envId=top-interview-150
func SpiralOrder() {
	var matrix [][]int
	matrix = append(matrix, []int{1, 2, 3, 4})
	matrix = append(matrix, []int{5, 6, 7, 8})
	matrix = append(matrix, []int{9, 10, 11, 12})
	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	var result []int
	if len(matrix) == 0 {
		return result
	}
	top, left, right, bottom := 0, 0, len(matrix[0])-1, len(matrix)-1
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}

	}
	return result
}

package interview150

import "fmt"

// Rotate Image
// https://leetcode.com/problems/rotate-image/description/?envType=study-plan-v2&envId=top-interview-150
func RotateImage() {
	var matrix [][]int
	matrix = append(matrix, []int{1, 2, 3})
	matrix = append(matrix, []int{4, 5, 6})
	matrix = append(matrix, []int{7, 8, 9})
	rotateImage(matrix)
}

func rotateImage(matrix [][]int) {
	top, left, bottom, right := 0, 0, len(matrix)-1, len(matrix)-1

	for left <= right && top <= bottom {
		for i := 0; i < right-left; i++ {
			tmp := matrix[top][left+i]
			matrix[top][left+i] = matrix[bottom-i][left]
			matrix[bottom-i][left] = matrix[bottom][right-i]
			matrix[bottom][right-i] = matrix[top+i][right]
			matrix[top+i][right] = tmp

		}
		top++
		left++
		right--
		bottom--
	}
	fmt.Println(matrix)
}

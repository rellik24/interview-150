package interview150

import "fmt"

// Set Matrix Zeroes
// https://leetcode.com/problems/set-matrix-zeroes/description/?envType=study-plan-v2&envId=top-interview-150
func SetZeroes() {
	var matrix [][]int
	matrix = append(matrix, []int{1, 1, 1})
	matrix = append(matrix, []int{1, 0, 1})
	matrix = append(matrix, []int{1, 1, 1})
	setZeroes(matrix)
}

func setZeroes(matrix [][]int) {
	var check [][]bool
	for _, v := range matrix {
		var tmp []bool
		for _, val := range v {
			if val == 0 {
				tmp = append(tmp, true)
			} else {
				tmp = append(tmp, false)
			}
		}
		check = append(check, tmp)
	}
	fmt.Println(check)

	for i, mtx := range matrix {
		for j, v := range mtx {
			if v == 0 && check[i][j] {

				fmt.Println(i, j)
				for idx := 0; idx < len(matrix); idx++ {
					matrix[idx][j] = 0
				}

				for idx := 0; idx < len(mtx); idx++ {
					matrix[i][idx] = 0
				}
			}
		}
	}
	fmt.Println(matrix)
}

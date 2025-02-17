package interview150

import "fmt"

// Valid Sudoku
// https://leetcode.com/problems/valid-sudoku/submissions/1507889938/?envType=study-plan-v2&envId=top-interview-150
// Note: 題目只要求確認是否合規，不管能不能真的解出來
func IsValidSudoku() {
	var board [][]byte
	board = append(board, []byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'})
	board = append(board, []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'})
	board = append(board, []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'})
	board = append(board, []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'})
	board = append(board, []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'})
	board = append(board, []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'})
	board = append(board, []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'})
	board = append(board, []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'})
	board = append(board, []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'})
	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {

	rows := [9][9]bool{}
	cols := [9][9]bool{}
	boxes := [9][9]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				v := board[i][j] - '1'
				box := (i/3)*3 + j/3
				if rows[i][v] || cols[j][v] || boxes[box][v] {
					return false
				} else {
					rows[i][v] = true
					cols[j][v] = true
					boxes[box][v] = true
				}
			}
		}
	}
	return true
}

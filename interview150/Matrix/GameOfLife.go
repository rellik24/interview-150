package interview150

// Game of Life
// https://leetcode.com/problems/game-of-life/description/?envType=study-plan-v2&envId=top-interview-150
func GameOfLife() {
	var matrix [][]int
	matrix = append(matrix, []int{0, 1, 0})
	matrix = append(matrix, []int{0, 0, 1})
	matrix = append(matrix, []int{1, 1, 1})
	matrix = append(matrix, []int{0, 0, 0})
	gameOfLife(matrix)
}

func gameOfLife(matrix [][]int) {
	var check [][]bool
	for _, v := range matrix {
		var tmp []bool
		for _, val := range v {
			if val == 1 { // live
				tmp = append(tmp, true)
			} else { // dead
				tmp = append(tmp, false)
			}
		}
		check = append(check, tmp)
	}
	// fmt.Println(check)

	for i, mtx := range matrix {
		for j, v := range mtx {
			neighbors := checkNeighbor(matrix, check, i, j)
			// fmt.Println(i, j, neighbors)
			if v == 1 && check[i][j] {
				if neighbors < 2 || neighbors > 3 {
					matrix[i][j] = 0
				}
			} else if v == 0 && !check[i][j] {
				if neighbors == 3 {
					matrix[i][j] = 1
				}
			}
		}
	}
	// fmt.Println(matrix)
}

// 檢查周邊數量
func checkNeighbor(matrix [][]int, check [][]bool, i, j int) int {
	neighbor := 0
	rows, cols := len(matrix), len(matrix[0])

	// 定義周邊 8 個方向的偏移量
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // 上左、上、上右
		{0, -1}, {0, 1}, // 左、中、右
		{1, -1}, {1, 0}, {1, 1}, // 下左、下、下右
	}
	for _, v := range directions {
		newi, newj := i+v[0], j+v[1]
		if newi < 0 || newj < 0 || newi >= rows || newj >= cols {

		} else {
			if check[newi][newj] {
				neighbor++
			}
		}
	}
	return neighbor
}

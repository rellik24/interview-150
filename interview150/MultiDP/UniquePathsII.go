package interview150

import "fmt"

func UniquePathsWithObstacles() {
	var grid [][]int
	// grid = append(grid, []int{1, 0})
	grid = append(grid, []int{0, 0, 0})
	grid = append(grid, []int{0, 1, 0})
	grid = append(grid, []int{0, 0, 0})
	fmt.Println(uniquePathsWithObstacles(grid))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	grid := make([][]int, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		tmp := make([]int, len(obstacleGrid[i]))
		grid[i] = tmp
	}

	if obstacleGrid[0][0] == 0 {
		grid[0][0] = 1
	}

	for i := 1; i < len(obstacleGrid); i++ {
		if obstacleGrid[i-1][0] == 0 && obstacleGrid[i][0] == 0 {
			grid[i][0] += grid[i-1][0]
		}
	}

	for i := 1; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i-1] == 0 && obstacleGrid[0][i] == 0 {
			grid[0][i] += grid[0][i-1]
		}
	}

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j-1]
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

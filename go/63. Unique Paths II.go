// 2/14
package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var i, j int

	// 障害物は-1で初期化
	for i = range obstacleGrid {
		for j = range obstacleGrid[i] {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = -1
			}
		}
	}
	if obstacleGrid[0][0] == -1 || obstacleGrid[i][j] == -1 {
		return 0
	}

	// 上端、左端を1で初期化
	for i = range obstacleGrid {
		if obstacleGrid[i][0] == -1 {
			break
		} else {
			obstacleGrid[i][0] = 1
		}
	}
	for j = range obstacleGrid[i] {
		if obstacleGrid[0][j] == -1 {
			break
		} else {
			obstacleGrid[0][j] = 1
		}
	}

	for i = range obstacleGrid {
		for j = range obstacleGrid[i] {
			if i == 0 || j == 0 {
				continue
			} else if obstacleGrid[i][j] == -1 {
				continue
			}
			if obstacleGrid[i-1][j] != -1 {
				obstacleGrid[i][j] += obstacleGrid[i-1][j]
			}
			if obstacleGrid[i][j-1] != -1 {
				obstacleGrid[i][j] += obstacleGrid[i][j-1]
			}
		}
	}

	return obstacleGrid[i][j]
}

type TestSet struct {
	Grid     [][]int
	Expected int
}

func main() {
	inputs := []TestSet{
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
		{[][]int{{0, 1}, {0, 0}}, 1},
		{[][]int{{0, 0}, {0, 1}}, 0},
		{[][]int{{0, 0}, {1, 1}, {0, 0}}, 0},
	}

	for _, v := range inputs {
		ret := uniquePathsWithObstacles(v.Grid)
		fmt.Println("Exp:", v.Expected, "Your Output:", ret)
	}

	return
}

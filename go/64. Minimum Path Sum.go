// 2/15
package main

import "fmt"

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[m-1][n-1]
}

type TestSet struct {
	Grid     [][]int
	Expected int
}

func main() {
	inputs := []TestSet{
		{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, 12},
		{[][]int{{1}}, 1},
		{[][]int{{1, 1}, {1, 1}}, 3},
	}

	for _, v := range inputs {
		ret := minPathSum(v.Grid)
		fmt.Println("Exp:", v.Expected, "Your Output:", ret)
	}

	return
}

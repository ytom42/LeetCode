// 2/12
package main

import "fmt"

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

type TestSet struct {
	M, N     int
	Expected int
}

func main() {
	inputs := []TestSet{
		{3, 7, 28},
		{3, 2, 3},
	}

	for _, v := range inputs {
		ret := uniquePaths(v.M, v.N)
		fmt.Println("Exp:", v.Expected, "Your Output:", ret)
	}

	return
}

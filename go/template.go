package main

import "fmt"

type TestSet struct {
	Expected int
}

func main() {
	inputs := []TestSet{}

	for _, v := range inputs {
		fmt.Println("Exp:", v.Expected, "Your Output:")
	}

	return
}

// ----------------------------------------------------------

func makeDP(m, n, init int) [][]int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = init
		}
	}

	return dp
}

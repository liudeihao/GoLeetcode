package main

import "fmt"

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = grid[m-1][n-1]
	for i := m - 2; i >= 0; i-- {
		dp[i][n-1] = grid[i][n-1] + dp[i+1][n-1]
	}
	for j := n - 2; j >= 0; j-- {
		dp[m-1][j] = grid[m-1][j] + dp[m-1][j+1]
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = grid[i][j] + min(dp[i+1][j], dp[i][j+1])
		}
	}
	for i := range m {
		fmt.Println(dp[i])
	}
	return dp[0][0]
}

package main

import (
	"slices"
)

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	dp := make([]int, m)
	dp[0] = triangle[0][0]
	for i := 1; i < m; i++ {
		dp[i] = dp[i-1] + triangle[i][i]
		// 注意这里是[i-1:1]
		for j := i - 1; j > 0; j-- {
			dp[j] = min(dp[j-1], dp[j]) + triangle[i][j]
		}
		dp[0] += triangle[i][0]
	}
	return slices.Min(dp)
}

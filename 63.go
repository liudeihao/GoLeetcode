package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	// 丑陋的边界情况
	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	} else if m == 1 && n == 1 {
		return 1
	}

	dp := make([]int, n)
	dp[n-1] = 1
	for i := m - 2; i >= 0; i-- {
		if obstacleGrid[i][n-1] == 1 {
			dp[n-1] = 0
		}
		for j := n - 2; j >= 0; j-- {
			if obstacleGrid[i][j] == 0 {
				dp[j] = dp[j] + dp[j+1]
			} else {
				dp[j] = 0
			}
		}
	}
	return dp[0]
}

package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	// 丑陋的边界情况
	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	} else if m == 1 && n == 1 {
		return 1
	}

	dp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, n)
	}

	for i := m - 2; i >= 0; i-- {
		if obstacleGrid[i][n-1] == 0 {
			dp[i][n-1] = 1
		} else {
			break
		}
	}
	for j := n - 2; j >= 0; j-- {
		if obstacleGrid[m-1][j] == 0 {
			dp[m-1][j] = 1
		} else {
			break
		}
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i+1][j] + dp[i][j+1]
			}
		}
	}
	return dp[0][0]
}

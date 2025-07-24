package main

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := i - 1; j >= 1; j-- {
			// 这个 j*(i-j) 是直接分成两个，因为dp[2]==1是把2分了两个，导致k>=3了
			dp[i] = max(dp[i], max(dp[j]*(i-j), j*(i-j)))
		}
	}
	return dp[n]
}

package main

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	dp := make([][2]int, n+1)
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}
	return max(dp[n-1][1], 0)
}

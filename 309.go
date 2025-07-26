package main

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 1 {
		return 0
	}
	dp := make([][3]int, n)
	// dp[i][0] 第i天持有股票
	// dp[i][1] 第i天卖出股票
	// dp[i][2] 第i天是冷冻期，即前一天卖出了股票
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		dp[i][2] = dp[i-1][1]
	}
	return max(dp[n-1][1], dp[n-1][2])
}

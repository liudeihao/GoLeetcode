package main

func maxProfit(prices []int) int {
	// 三种状态：
	// 0. 持有股票，之前没卖出过
	// 1. 不持有股票，之前没卖出过
	// 2. 持有股票，之前卖出过
	// 3. 不持有股票，之前卖出过
	n := len(prices)
	dp := make([][4]int, n+1)
	dp[0][0] = -prices[0]
	dp[0][2] = -prices[0] // 相当于第一天当天买入当天卖出，然后买入...
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]-prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]+prices[i])
	}
	return max(dp[n-1][1], dp[n-1][3])
}

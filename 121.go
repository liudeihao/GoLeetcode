package main

// 动态规划
func maxProfit(prices []int) int {
	dp := make([][2]int, len(prices)+1)
	// dp[i][0]为第i天持有股票
	// dp[i][1]为第i天不持有股票
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}

/*
// 贪心
func maxProfit_121(prices []int) int {
	profit := 0
	lowest := 100_000
	for _, price := range prices {
		lowest = min(lowest, price)
		profit = max(profit, price-lowest)
	}
	return profit
}
*/

package main

// 动态规划
func maxProfit(prices []int) int {
	dp := make([][2]int, len(prices)+1)
	// dp[i][0]为第i天持有股票
	// dp[i][1]为第i天不持有股票
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}

/*
// 贪心
func maxProfit_122(prices []int) int {
	pre := prices[0]
	ret := 0
	for _, cur := range prices {
		if cur > pre {
			ret += cur - pre
		}
		pre = cur
	}
	return ret
}
*/

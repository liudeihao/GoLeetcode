package main

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, c := range coins {
		for j := c; j <= amount; j++ {
			dp[j] += dp[j-c] // 求组合数
		}
	}
	return dp[amount]
}

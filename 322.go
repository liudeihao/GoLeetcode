package main

func coinChange(coins []int, amount int) int {
	dp := make([]int, 10001)
	Max := amount + 1
	for i := 1; i <= amount; i++ {
		dp[i] = Max
	}
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i >= c {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}
	if dp[amount] < Max {
		return dp[amount]
	}
	return -1
}

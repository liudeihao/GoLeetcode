package main

func numSquares(n int) int {
	var nums []int
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = n
	}
	for i := 1; i*i <= n; i++ {
		nums = append(nums, i*i)
		dp[i*i] = 1
	}
	for _, v := range nums {
		for j := v; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-v]+1)
		}
	}
	return dp[n]
}

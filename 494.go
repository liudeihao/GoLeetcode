package main

func findTargetSumWays(nums []int, target int) int {
	// 和 1049.最后一块石头的重量 II 一样...
	sum := 0
	for _, v := range nums {
		sum += v
	}
	diff := sum - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	neg := diff / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	for _, num := range nums {
		for j := neg; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[neg]
}

/*
一开始的解法

func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	dp := make([]int, 2005)
	dpPrev := make([]int, 2005)
	dpPrev[1000+nums[0]]++
	dpPrev[1000-nums[0]]++
	for i := 1; i < n; i++ {
		for j := 1000 - nums[i]; j >= -1000+nums[i]; j-- {
			dp[1000+j+nums[i]] += dpPrev[1000+j]
			dp[1000+j-nums[i]] += dpPrev[1000+j]
		}
		dpPrev = dp
		dp = make([]int, 2005)
	}
	return dpPrev[1000+target]
}
*/

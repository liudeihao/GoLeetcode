package main

func lastStoneWeightII(stone []int) int {
	// 让两个子集和尽量接近
	n := len(stone)
	sum := 0
	for _, num := range stone {
		sum += num
	}
	dp := make([]int, sum/2+1)
	for i := 0; i < n; i++ {
		for j := sum / 2; j >= stone[i]; j-- {
			dp[j] = max(dp[j], dp[j-stone[i]]+stone[i])
		}
	}
	return sum - 2*dp[sum/2]
}

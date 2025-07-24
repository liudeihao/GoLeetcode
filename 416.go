package main

func canPartition(nums []int) bool {
	// 其实就是找有没有子集和==sum/2
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	dp := make([]bool, sum+1)
	dp[nums[0]] = true
	for i := 1; i < n; i++ {
		// dp[j]表示前i个值是否有子集和为j
		for j := sum / 2; j >= nums[i]; j-- { // 只需要前sum/2个
			dp[j] = dp[j] || dp[j-nums[i]]
		}
		dp[nums[i]] = true
	}
	return dp[sum/2]
}

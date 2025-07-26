package main

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	// 爬楼梯，nums是可能的跳跃距离
	for j := 1; j <= target; j++ {
		for _, v := range nums {
			if v <= j {
				dp[j] += dp[j-v]
			}
		}
	}
	return dp[target]
}

/*
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for j := 1; j <= target; j++ {
		// 如果求组合数就是外层for循环遍历物品，内层for遍历背包。
		// 如果求排列数就是外层for遍历背包，内层for循环遍历物品。
		for _, v := range nums {
			if j >= v {
				dp[j] += dp[j-v]
			}
		}
	}
	return dp[target]
}
*/

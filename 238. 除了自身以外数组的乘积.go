package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prodL, prodR := make([]int, n), make([]int, n)
	prodL[0] = 1
	prodR[n-1] = 1
	for i := 1; i < n; i++ {
		// prodL[i] 定义为 i 左边（不包括i）的连乘
		prodL[i] = prodL[i-1] * nums[i-1]
		prodR[n-i-1] = prodR[n-i] * nums[n-i]
	}
	ans := make([]int, n)
	ans[0] = prodR[0]
	ans[n-1] = prodL[n-1]
	for i := 1; i < n-1; i++ {
		// ans[i] 就是之前定义这两个数组【对应位置】的乘积即可。
		// 依赖这个所谓【对应位置】，就不容易把自己绕晕。
		ans[i] = prodL[i] * prodR[i]
	}
	return ans
}

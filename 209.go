package main

import "math"

func minSubArrayLen(target int, nums []int) int {
	// 滑动窗口。
	/*
		首先要思考 如果用一个for循环，那么应该表示 滑动窗口的起始位置，还是终止位置。
		如果只用一个for循环来表示 滑动窗口的起始位置，那么如何遍历剩下的终止位置？
		此时难免再次陷入 暴力解法的怪圈。
		所以 只用一个for循环，那么这个循环的索引，一定是表示 滑动窗口的终止位置。
	*/
	n := len(nums)
	sum, ans := 0, math.MaxInt
	l := 0
	for r := 0; r < n; r++ { // 遍历右指针
		sum += nums[r]
		for sum >= target { // 这里是滑动窗口的核心
			ans = min(ans, r-l+1)
			sum -= nums[l]
			l++
		}
	}
	if ans == math.MaxInt {
		return 0
	} else {
		return ans
	}
}

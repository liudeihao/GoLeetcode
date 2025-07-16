package main

import "slices"

func sortedSquares(nums []int) []int {
	n := len(nums)
	var ans []int
	l, r := 0, n-1
	for l <= r {
		ls, rs := nums[l]*nums[l], nums[r]*nums[r]
		if ls > rs {
			ans = append(ans, ls)
			l++
		} else {
			ans = append(ans, rs)
			r--
		}
	}
	slices.Reverse(ans)
	return ans
}

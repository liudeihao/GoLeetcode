package main

import "sort"

func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	var first int
	if nums[0] < 0 {
		first = sort.Search(n, func(i int) bool {
			return nums[i] > 0
		})
	}
	neg, pos := first-1, first
	p := 0
	for neg >= 0 && pos <= n-1 {
		if nums[neg]*nums[neg] < nums[pos]*nums[pos] {
			ans[p] = nums[neg] * nums[neg]
			neg--
		} else {
			ans[p] = nums[pos] * nums[pos]
			pos++
		}
		p++
	}
	for neg >= 0 {
		ans[p] = nums[neg] * nums[neg]
		neg--
		p++
	}
	for pos <= n-1 {
		ans[p] = nums[pos] * nums[pos]
		pos++
		p++
	}
	return ans
}

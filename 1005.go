package main

import "slices"

func largestSumAfterKNegations(nums []int, k int) int {
	slices.Sort(nums)
	n := len(nums)
	ans, i := 0, 0

	for i < n && nums[i] < 0 && k > 0 {
		ans += -nums[i]
		k--
		i++
	}
	for j := i; j < n; j++ {
		ans += nums[j]
	}
	if k%2 == 1 {
		var minPos int
		if i == 0 {
			minPos = nums[0]
		} else if i == n {
			minPos = nums[n-1]
		} else {
			minPos = min(-nums[i-1], nums[i])
		}
		ans -= 2 * minPos
	}
	return ans
}

package main

import "fmt"

func summaryRanges(nums []int) []string {
	var ans []string
	l, r, n := 0, 1, len(nums)
	for l < n {
		if r < n && nums[r] == nums[r-1]+1 {
			r++
			continue
		}
		if nums[r-1] == nums[l] {
			ans = append(ans, fmt.Sprintf("%d", nums[l]))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", nums[l], nums[r-1]))
		}
		l = r
		r++
	}
	return ans
}

package main

func searchInsert(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n
	for l < r {
		m := (l + r) / 2
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

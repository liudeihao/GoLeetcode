package main

func search_704(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (l + r) / 2
		if nums[m] > target {
			r = m
		} else if nums[m] == target {
			return m
		} else {
			l = m + 1
		}
	}
	return -1
}

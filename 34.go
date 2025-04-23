package main

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	search := func(n int, condition func(int) bool) int {
		l, r := 0, n
		for l < r {
			m := (l + r) / 2
			if condition(m) {
				r = m
			} else {
				l = m + 1
			}
		}
		return l
	}
	first := search(n, func(idx int) bool {
		return nums[idx] >= target
	})
	last := search(n, func(idx int) bool {
		return nums[idx] > target
	}) - 1
	if first >= n || nums[first] != target {
		return []int{-1, -1}
	}
	return []int{first, last}
}

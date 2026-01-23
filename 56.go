package main

import (
	"cmp"
	"slices"
)

func merge_56(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(x, y []int) int {
		if x[0] == y[0] {
			return cmp.Compare(x[1], y[1])
		}
		return cmp.Compare(x[0], y[0])
	})
	var ans [][]int

	left, right := intervals[0][0], intervals[0][1]
	for _, interval := range intervals {
		if interval[0] <= right {
			right = max(interval[1], right)
		} else {
			ans = append(ans, []int{left, right})
			left, right = interval[0], interval[1]
		}
	}
	ans = append(ans, []int{left, right})
	return ans
}

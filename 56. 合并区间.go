package main

import "slices"

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	var ans [][]int
	prevL, prevR := intervals[0][0], intervals[0][1]
	for _, interval := range intervals {
		l, r := interval[0], interval[1]
		if l > prevR {
			ans = append(ans, []int{prevL, prevR})
			prevL = l
		}
		prevR = max(prevR, r)
	}
	return append(ans, []int{prevL, prevR})
}

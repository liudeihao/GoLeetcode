package main

import (
	"cmp"
	"slices"
)

func findMinArrowShots(points [][]int) int {
	slices.SortFunc(points, func(x, y []int) int {
		if x[0] == y[0] {
			return cmp.Compare(x[1], y[1])
		}
		return cmp.Compare(x[0], y[0])
	})
	ans := 0
	r := len(points) - 1
	for r >= 0 {
		l := points[r][0]
		for r >= 0 && points[r][1] >= l {
			r--
		}
		ans++
	}
	return ans
}

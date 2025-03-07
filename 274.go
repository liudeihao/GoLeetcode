package main

import "slices"

func hIndex(citations []int) int {
	slices.SortFunc(citations, func(x, y int) int {
		return y - x
	})
	c := 0
	for _, v := range citations {
		if v <= c {
			break
		}
		c++
	}
	return c
}

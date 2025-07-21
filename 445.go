package main

import "slices"

func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)
	var ans int

	for b, c := len(g)-1, len(s)-1; b >= 0 && c >= 0; b-- {
		if g[b] <= s[c] {
			c--
			ans++
		}
	}
	return ans
}

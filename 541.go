package main

import "slices"

func reverseStr(s string, k int) string {
	n := len(s)
	ans := []byte(s)
	for i := 0; i < n; i += 2 * k {
		if i+k < n {
			slices.Reverse(ans[i : i+k])
		} else {
			slices.Reverse(ans[i:n])
		}
	}
	return string(ans)
}

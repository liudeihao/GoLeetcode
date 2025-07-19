package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	isIP := func(s string) bool {
		if len(s) > 1 && s[0] == '0' {
			return false
		}
		i, err := strconv.Atoi(s)
		if err == nil && i >= 0 && i <= 255 {
			return true
		}
		return false
	}

	n := len(s)
	var ans []string
	var path []string
	var backtrack func(startIdx int)
	backtrack = func(startIdx int) {
		if len(path) == 4 && startIdx >= n {
			ans = append(ans, strings.Join(path, "."))
			return
		}
		for endIdx := startIdx + 1; endIdx <= n; endIdx++ {
			if !isIP(s[startIdx:endIdx]) {
				break
			}
			path = append(path, s[startIdx:endIdx])
			backtrack(endIdx)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return ans
}

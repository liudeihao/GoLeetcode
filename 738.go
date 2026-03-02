package main

import "strconv"

func monotoneIncreasingDigits(n int) int {
	var ds []byte
	ds = []byte(strconv.Itoa(n))
	if len(ds) == 1 {
		return n
	}
	allNine := len(ds)
	// 从后遍历，找到 高位>低维 的
	for j := len(ds) - 1; j > 0; j-- {
		if ds[j-1] > ds[j] {
			allNine = j // 永远在借位的那里
			ds[j-1]--   // 让 高位-- 后继续往高位找 高位>低维
		}
	}
	for allNine < len(ds) {
		ds[allNine] = '9'
		allNine++
	}
	ans, _ := strconv.Atoi(string(ds))
	return ans
}

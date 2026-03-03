package main

func countBinarySubstrings(s string) int {
	ans := 0
	prev, curr := 0, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			curr++
		} else {
			ans += min(prev, curr)
			prev, curr = curr, 1
		}
	}
	ans += min(prev, curr)
	return ans
}

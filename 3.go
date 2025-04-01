package main

func lengthOfLongestSubstring(s string) int {
	set := make(map[byte]bool)
	l := 0
	ans := 0
	for r := range s {
		rch := s[r]
		for set[rch] && l < r {
			lch := s[l]
			set[lch] = false
			l++
		}
		set[rch] = true
		ans = max(ans, r-l+1)
	}
	return ans
}

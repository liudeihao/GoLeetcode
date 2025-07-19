package main

func partition(s string) [][]string {
	isPal := func(s string) bool {
		l, r := 0, len(s)-1
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}
	
	n := len(s)
	var ans [][]string
	var path []string
	var backtrack func(startIdx int)
	backtrack = func(startIdx int) {
		if startIdx >= n {
			tmp := make([]string, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for endIdx := startIdx + 1; endIdx <= n; endIdx++ {
			if !isPal(s[startIdx:endIdx]) {
				continue
			}
			path = append(path, s[startIdx:endIdx])
			backtrack(endIdx)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return ans
}

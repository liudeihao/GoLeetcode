package main

func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n+1)
	for l := 0; l <= n; l++ {
		dp[l] = make([]bool, n+1)
		dp[l][0] = true
	}
	var ans int
	for l := n - 1; l >= 0; l-- {
		for r := l; r < n; r++ {
			if s[l] == s[r] && (r-l <= 1 || dp[l+1][r-1]) {
				dp[l][r] = true
				ans++
			}
		}
	}
	return ans
}

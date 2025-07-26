package main

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for j := 1; j <= n; j++ {
		for _, word := range wordDict {
			if j >= len(word) && word == s[j-len(word):j] {
				dp[j] = dp[j] || dp[j-len(word)]
			}
		}
	}
	return dp[n]
}

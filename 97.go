package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n1+n2 != n3 {
		return false
	}
	// 滚动数组优化
	dp := make([]bool, n2+1)
	dp[0] = true
	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			p := i + j - 1
			if i > 0 {
				dp[j] = dp[j] && s1[i-1] == s3[p]
			}
			if j > 0 {
				dp[j] = dp[j] || dp[j-1] && s2[j-1] == s3[p]
			}
		}
	}
	return dp[n2]
}

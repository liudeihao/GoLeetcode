package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n1+n2 != n3 {
		return false
	}
	// 定义f(i,j)表示s1[:i]与s2[:j]可以交错组成s3[:i+j]
	dp := make([][]bool, n1+1)
	for i := range dp {
		dp[i] = make([]bool, n2+1)
	}
	dp[0][0] = true
	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}
	return dp[n1][n2]
}

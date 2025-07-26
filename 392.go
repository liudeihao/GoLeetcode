package main

// 动态规划
func isSubsequence(s string, t string) bool {
	// 问的是s是不是t的子序列
	n1, n2 := len(s), len(t)
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			// dp[i][j]: s[:i]和t[:j]相同子序列的长度。
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[n1][n2] == n1
}

/*
// 双指针
func isSubsequence(s string, t string) bool {
	ns, nt := len(s), len(t)
	ps, pt := 0, 0
	for ps < ns && pt < nt {
		if s[ps] == t[pt] {
			ps++
			pt++
		} else {
			pt++
		}
	}
	return ps == ns
}
*/

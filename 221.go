package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	ans := 0
	dp := [300][300]int{}
	for i := 0; i < m; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			ans = 1
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			ans = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
				continue
			}
			lu := dp[i-1][j-1]
			u, l := dp[i-1][j], dp[i][j-1]
			if u >= lu && l >= lu {
				dp[i][j] = lu + 1
			} else {
				dp[i][j] = min(u, l) + 1
			}
			ans = max(ans, dp[i][j])
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", dp[i][j])
		}
		fmt.Println()
	}
	return ans * ans
}

func main() {
	n := 300
	a := make([][]byte, 300)
	for i := 0; i < n; i++ {
		a[i] = make([]byte, 300)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[i][j] = '1'
		}
	}
	maximalSquare(a)

}

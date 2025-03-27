package main

func climbStairs(n int) int {
	prev1, prev2 := 1, 0
	ans := 0
	for i := 1; i <= n; i++ {
		ans = prev1 + prev2
		prev2 = prev1
		prev1 = ans
	}
	return ans
}

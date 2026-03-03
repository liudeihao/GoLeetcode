package main

func concatenatedBinary(n int) int {
	mod := 1e9 + 7
	shift := 0
	ans := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			shift++
		}
		ans = (ans<<shift + i) % int(mod)
	}
	return ans
}

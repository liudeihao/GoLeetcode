package main

func binaryGap(n int) int {
	ans := 0
	prev := -1
	p := 0
	for (n >> p) > 0 {
		if (n>>p)&1 == 1 {
			if prev != -1 {
				ans = max(ans, p-prev)
			}
			prev = p
		}
		p++
	}
	return ans
}

package main

func totalMoney(n int) int {
	monday := 1
	ans := 0
	for i := 0; i < n; i++ {
		ans += monday + i%7
		if i%7 == 6 {
			monday++
		}
	}
	return ans
}

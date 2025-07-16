package main

func isHappy(n int) bool {
	visit := make(map[int]bool)
	for n != 1 {
		if visit[n] {
			return false
		}
		visit[n] = true
		next := 0
		for n > 0 {
			r := n % 10
			n = n / 10
			next += r * r
		}
		n = next
	}
	return true
}

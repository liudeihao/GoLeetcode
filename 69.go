package main

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	l, r := 0, x
	for l < r {
		m := l + (r-l)/2
		if m*m > x {
			r = m
		} else {
			l = m + 1
		}
	}
	return r - 1
}

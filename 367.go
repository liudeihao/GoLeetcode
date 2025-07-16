package main

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	l, r := 0, num
	for l < r {
		m := l + (r-l)/2
		if m*m > num {
			r = m
		} else if m*m == num {
			return true
		} else {
			l = m + 1
		}
	}
	return false
}

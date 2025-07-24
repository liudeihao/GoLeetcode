package main

func fib(n int) int {
	if n <= 1 {
		return n
	}
	pre1 := 0
	pre2 := 1
	for n-1 > 0 {
		n--
		pre1, pre2 = pre2, pre1+pre2
	}
	return pre2
}

package main

func smallestNumber(n int) int {
	// 完全就是求highbit的前置。
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	return n
}

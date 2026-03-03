package main

import "math/bits"

func countPrimeSetBits(left int, right int) int {
	ans := 0
	// 事实上primes可以直接用二进制数表示，后续用移位判断即可。
	primes := make(map[int]bool)
	for _, v := range []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31} {
		primes[v] = true
	}
	for i := left; i <= right; i++ {
		if primes[bits.OnesCount(uint(i))] {
			ans++
		}
	}
	return ans
}

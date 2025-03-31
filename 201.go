package main

func rangeBitwiseAnd(left int, right int) int {
	if left == right {
		return right
	} else {
		return 2 * rangeBitwiseAnd(left>>1, right>>1)
	}
}

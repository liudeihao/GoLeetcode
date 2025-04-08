package main

func singleNumber(nums []int) int {
	a, b := 0, 0
	for _, num := range nums {
		// 以下优化自 a, b = b&^a&num|a&^b&^num, (b^num)&^a
		b = (b ^ num) &^ a
		a = (a ^ num) &^ b
	}
	return b
}

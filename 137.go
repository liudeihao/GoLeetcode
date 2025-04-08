package main

func singleNumber(nums []int) int {
	// 答案的第i个二进制位就是数组中所有元素的第i个二进制位之和除以3的余数
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}

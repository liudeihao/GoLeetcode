package main

// 但是题目要求原地算法
func majorityElement(nums []int) int {
	n := len(nums) / 2
	mp := make(map[int]int)

	for _, num := range nums {
		mp[num]++
	}
	for k, v := range mp {
		if v > n {
			return k
		}
	}
	return -1
}

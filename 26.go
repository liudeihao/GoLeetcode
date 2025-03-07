package main

func removeDuplicates(nums []int) int {
	// 双指针
	n := len(nums)
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

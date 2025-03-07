package main

func removeElement(nums []int, val int) int {
	// 双指针
	n := len(nums)
	slow := 0
	for fast := 0; fast < n; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

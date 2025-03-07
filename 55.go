package main

func canJump(nums []int) bool {
	n := len(nums)
	prev := nums[0]
	for i := 1; i < n; i++ {
		if prev < i {
			return false
		}
		prev = max(prev, i+nums[i])
	}
	return true
}

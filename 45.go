package main

func jump(nums []int) int {
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, nums[i]+i)
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	_rob := func(nums []int) int {
		dpVisit, dpSkip := 0, 0
		for _, v := range nums {
			dpSkip, dpVisit = dpVisit, max(dpVisit, dpSkip+v)
		}
		return max(dpSkip, dpVisit)
	}
	return max(_rob(nums[:len(nums)-1]), _rob(nums[1:]))
}

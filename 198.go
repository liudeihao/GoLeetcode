package main

func rob(nums []int) int {
	dpVisit, dpSkip := 0, 0
	for _, v := range nums {
		dpSkip, dpVisit = dpVisit, max(dpVisit, dpSkip+v)
	}
	return max(dpSkip, dpVisit)
}

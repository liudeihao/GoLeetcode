package main

func maxIncreasingSubarrays(nums []int) int {
	ans := 0
	curr, prev := 1, 0
	for i := 0; i+1 < len(nums); i++ {
		if nums[i] >= nums[i+1] {
			prev, curr = curr, 1
		} else {
			curr++
		}
		ans = max(ans, min(curr, prev))
		ans = max(ans, curr/2)
	}
	return ans
}

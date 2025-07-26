package main

func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	ans, cnt := 1, 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			cnt++
			ans = max(cnt, ans)
		} else {
			cnt = 1
		}
	}
	return max(cnt, ans)
}

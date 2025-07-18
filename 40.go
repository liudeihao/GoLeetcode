package main

import "slices"

func combinationSum2(nums []int, target int) [][]int {
	slices.Sort(nums)

	var ans [][]int
	var path []int
	var sum int

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for j := idx; j < len(nums) && nums[j]+sum <= target; j++ {
			if j > idx && nums[j] == nums[j-1] {
				// 只需要第一个idx就够了，它后面会遍历到后面的。
				// 多想想
				continue
			}
			path = append(path, nums[j])
			sum += nums[j]
			backtrack(j + 1)
			sum -= nums[j]
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return ans
}

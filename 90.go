package main

import "slices"

func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	n := len(nums)
	var ans [][]int
	var path []int
	var backtrack func(i int)
	backtrack = func(i int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		ans = append(ans, tmp)

		for j := i; j < n; j++ {
			if j != i && nums[j] == nums[j-1] {
				continue
			}
			path = append(path, nums[j])
			backtrack(j + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return ans
}

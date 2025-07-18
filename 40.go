package main

import "slices"

func combinationSum2(nums []int, target int) [][]int {
	slices.Sort(nums)

	var ans [][]int
	var path []int
	var sum int

	// 要去重的是同一层（兄弟间）的“使用过”，同一树枝上的都是一个组合里的元素，不用去重
	used := make([]bool, len(nums))

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for j := idx; j < len(nums) && nums[j]+sum <= target; j++ {
			if j > 0 && nums[j] == nums[j-1] && used[j-1] == false {
				// used[j-1] == false说明左兄弟用了nums[j-1]，去除。
				// used[j-1] == true说明父亲用了nums[j-1]，保留。
				continue
			}
			path = append(path, nums[j])
			sum += nums[j]
			used[j] = true
			backtrack(j + 1)
			used[j] = false
			sum -= nums[j]
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return ans
}

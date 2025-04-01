package main

import "slices"

func combinationSum(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	var ans [][]int
	var path []int
	var sum int
	var backtrack func(int)
	backtrack = func(start int) {
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if sum+candidates[i] > target {
				break
			}
			path = append(path, candidates[i])
			sum += candidates[i]
			backtrack(i)
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return ans
}

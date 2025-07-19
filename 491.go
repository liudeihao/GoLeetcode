package main

func findSubsequences(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	var path []int

	var backtrack func(int)
	backtrack = func(idx int) {
		if len(path) >= 2 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			// 子集问题，每个节点都可能是结果，而不仅仅是叶子。
			// 所以不用return
		}
		used := make(map[int]bool)
		// 用map，保证它的子节点中，重复的跳过。
		for j := idx; j < n; j++ {
			if used[nums[j]] {
				continue
			}
			if len(path) != 0 && nums[j] < path[len(path)-1] {
				continue
			}

			path = append(path, nums[j])
			used[nums[j]] = true
			backtrack(j + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return ans
}

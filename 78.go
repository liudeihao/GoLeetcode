package main

func subsets(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	var path []int
	var backtrack func(int)
	backtrack = func(startIdx int) {
		if startIdx >= n {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		backtrack(startIdx + 1)
		path = append(path, nums[startIdx])
		backtrack(startIdx + 1)
		path = path[:len(path)-1]

	}
	backtrack(0)
	return ans
}

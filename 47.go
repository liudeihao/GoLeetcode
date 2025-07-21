package main

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	var path []int
	vs := make(map[int]bool, n)

	var backtrack func()
	backtrack = func() {
		if len(path) == n {
			tmp := make([]int, n)
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		valVs := make(map[int]bool, n)
		for i, v := range nums {
			if valVs[v] || vs[i] {
				continue
			}
			path = append(path, nums[i])
			vs[i] = true
			valVs[nums[i]] = true
			backtrack()
			vs[i] = false
			path = path[:len(path)-1]
		}
	}
	backtrack()
	return ans

}

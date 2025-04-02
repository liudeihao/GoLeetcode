package main

import "fmt"

func permute(nums []int) [][]int {
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
		for i := range nums {
			if vs[nums[i]] {
				continue
			}
			path = append(path, nums[i])
			vs[nums[i]] = true
			backtrack()
			vs[nums[i]] = false
			path = path[:len(path)-1]
		}
	}
	backtrack()
	return ans
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

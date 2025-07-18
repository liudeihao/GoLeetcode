package main

func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var path []int
	var sum int
	var backtrack func(i int)
	backtrack = func(i int) {
		if sum == n && len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for j := i; j <= 9; j++ {
			if sum+j > n {
				break
			}
			path = append(path, j)
			sum += j
			backtrack(j)
			sum -= j
			path = path[:len(path)-1]
		}
	}
	backtrack(1)
	return ans
}

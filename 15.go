package main

import "slices"

func threeSum(nums []int) [][]int {
	// 哈希并不合适，要用双指针。
	var ans [][]int
	slices.Sort(nums)
	for i, n := range nums {
		if i > 0 && n == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			nl, nr := nums[l], nums[r]
			if nl+n+nr == 0 {
				ans = append(ans, []int{nl, n, nr})
				for l < r && nums[l] == nl {
					l++
				}
				for l < r && nums[r] == nr {
					r--
				}
			} else if nl+n+nr > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return ans
}

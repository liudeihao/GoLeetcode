package main

import "slices"

func fourSum(nums []int, target int) [][]int {
	var ans [][]int
	slices.Sort(nums)
	for i, n1 := range nums {
		if i != 0 && nums[i-1] == n1 {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			n2 := nums[j]
			if j != i+1 && nums[j-1] == n2 {
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				n3, n4 := nums[l], nums[r]
				sum := n1 + n2 + n3 + n4
				if sum == target {
					ans = append(ans, []int{n1, n2, n3, n4})
					for l < r && nums[l] == n3 {
						l++
					}
					for l < r && nums[r] == n4 {
						r--
					}
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return ans
}

package main

import "slices"

func threeSum(nums []int) [][]int {
    n := len(nums)
    slices.Sort(nums)
    var ans [][]int
    for i := 0; i < n-2; i++ {
        if i > 0 && nums[i] == nums[i-1] { // 去重
            continue
        }
        target := -nums[i]
        l, r := i+1, n-1

        for l < r {
            sum := nums[l] + nums[r]
            if sum == target {
                ans = append(ans, []int{nums[i], nums[l], nums[r]})
                l++ // 这两个数用了，指针都要移动
                r--
                for l < r && nums[l] == nums[l-1] { // 去重
                    l++
                }
                for l < r && nums[r] == nums[r+1] {
                    r--
                }
            } else if sum < target {
                l++
            } else {
                r--
            }
        }
    }
    return ans
}

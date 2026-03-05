package main

func moveZeroes(nums []int) {
    n := len(nums)
    l, r := 0, 0
    for r < n {
        for r < n && nums[r] == 0 {
            r++
        }
        if r < n {
            nums[l] = nums[r]
            r++
            l++
        }
    }
    for l < n {
        nums[l] = 0
        l++
    }
}

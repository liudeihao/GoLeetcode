package main

func maxArea(height []int) int {
    n := len(height)
    l, r := 0, n-1
    ans := 0
    for l < r {
        ans = max(ans, (r-l)*min(height[l], height[r]))
        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }
    return ans
}

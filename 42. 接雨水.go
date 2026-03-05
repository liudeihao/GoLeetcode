package main

func trap(height []int) int {
    n := len(height)
    leftMax, rightMax := make([]int, n), make([]int, n)
    mx := 0
    for i := 0; i < n; i++ {
        leftMax[i] = mx
        mx = max(mx, height[i])
    }
    mx = 0
    for i := n - 1; i >= 0; i-- {
        rightMax[i] = mx
        mx = max(mx, height[i])
    }
    var ans int
    for i := 0; i < n; i++ {
        ans += max(0, min(leftMax[i], rightMax[i])-height[i])
    }
    return ans
}

/*
没看题，写了一段计算最大接雨水量的代码，结果发现题目是求接雨水的总时间，才意识到自己理解错了题。

func trap(height []int) int {
    n := len(height)
    leftMaxIdx, rightMaxIdx := make([]int, n), make([]int, n)
    leftSum := make([]int, n)
    mx, mxIdx := 0, 0
    for i := 0; i < n; i++ {
        leftMaxIdx[i] = mxIdx
        if height[i] > mx {
            mxIdx = i
            mx = height[i]
        }
        leftSum[i] = height[i] + leftSum[max(0, i-1)]
    }
    mx, mxIdx = 0, n-1
    for i := n - 1; i >= 0; i-- {
        rightMaxIdx[i] = mxIdx
        if height[i] > mx {
            mxIdx = i
            mx = height[i]
        }
    }
    var ans int
    for i := 1; i < n-1; i++ {
        l, r := leftMaxIdx[i], rightMaxIdx[i]
        lv, rv := height[l], height[r]
        ans = max(ans, min(lv, rv)*(r-l-1)-(leftSum[r-1]-leftSum[l]))
    }
    return ans
}

*/

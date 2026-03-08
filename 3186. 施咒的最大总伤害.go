package main

import "slices"

func maximumTotalDamage(power []int) int64 {
    mp := make(map[int]int)
    for _, p := range power {
        mp[p]++
    }

    pc := make([][2]int, 0)

    for p, c := range mp {
        pc = append(pc, [2]int{p, c})
    }
    slices.SortFunc(pc, func(x, y [2]int) int {
        return x[0] - y[0]
    })

    n := len(pc)
    dp := make([]int, n)
    ans := 0

    // 不需要看很远的前面，最多看 2 个。因为如果 p-pc[j][0] <= 2，那么 dp[i] 就不可能从 j 往前转移了。
    j := 0 // 相当于慢指针
    prevSum := 0
    for i := 0; i < n; i++ {
        p, c := pc[i][0], pc[i][1]
        for j < i && p-pc[j][0] > 2 {
            if dp[j] > prevSum {
                prevSum = dp[j]
            }
            j++
        }
        dp[i] = prevSum + p*c
        ans = max(ans, dp[i])
    }
    return int64(ans)
}

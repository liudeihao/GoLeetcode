package main

import "math"

func maximumEnergy(energy []int, k int) int {
    ans := math.MinInt
    absorb := make([]int, k)
    for i := len(energy) - 1; i >= 0; i-- {
        absorb[i%k] += energy[i]
        ans = max(ans, absorb[i%k])
    }
    return ans
}

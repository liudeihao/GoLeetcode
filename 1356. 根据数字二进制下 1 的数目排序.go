package main

import "slices"

func sortByBits(arr []int) []int {
    countBits := func(x int) int {
        ans := 0
        for x > 0 {
            ans += x & 1
            x >>= 1
        }
        return ans
    }

    slices.SortFunc(arr, func(i, j int) int {
        c1, c2 := countBits(i), countBits(j)
        if c1 == c2 {
            return i - j
        }
        return c1 - c2
    })
    return arr
}

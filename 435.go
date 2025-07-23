package main

import (
    "cmp"
    "fmt"
    "math"
    "slices"
)

func eraseOverlapIntervals(intervals [][]int) int {
    slices.SortFunc(intervals, func(x, y []int) int {
        if x[0] == y[0] {
            return cmp.Compare(x[1], y[1])
        }
        return cmp.Compare(x[0], y[0])
    })
    ans := 0
    curR := math.MinInt
    for _, interval := range intervals {
        fmt.Println(curR, interval, interval[0] >= curR)
        if interval[0] >= curR {
            curR = interval[1]
        } else {
            ans++
            curR = min(curR, interval[1])
        }
    }
    return ans
}

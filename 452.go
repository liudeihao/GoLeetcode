package main

import (
    "cmp"
    "math"
    "slices"
)

func findMinArrowShots(points [][]int) int {
    slices.SortFunc(points, func(x, y []int) int {
        if x[0] == y[0] {
            return cmp.Compare(x[1], y[1])
        }
        return cmp.Compare(x[0], y[0])
    })
    ans := 1
    minR := math.MaxInt
    for _, p := range points {
        if p[0] <= minR {
            minR = min(p[1], minR)
            continue
        }
        minR = p[1]
        ans++
    }
    return ans
}

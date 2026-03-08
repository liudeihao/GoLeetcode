package main

import "sort"

func avoidFlood(rains []int) []int {
    n := len(rains)
    ans := make([]int, n)
    for i := 0; i < n; i++ {
        ans[i] = 1
    }
    sunny := make([]int, 0) // 可以抽干的日子

    lake := make(map[int]int)
    for i, rain := range rains {
        if rain == 0 {
            sunny = append(sunny, i)
            continue
        }
        ans[i] = -1
        if day, ok := lake[rain]; ok { // 这个湖已经有水了
            l := sort.SearchInts(sunny, day)

            if l == len(sunny) {
                return []int{}
            }
            ans[sunny[l]] = rain
            sunny = append(sunny[:l], sunny[l+1:]...)
        }
        lake[rain] = i
    }
    return ans
}

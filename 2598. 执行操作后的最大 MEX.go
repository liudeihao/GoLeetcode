package main

func findSmallestInteger(nums []int, value int) int {
    cnt := make([]int, value)
    for _, v := range nums {
        v %= value
        if v < 0 {
            v += value
        }
        cnt[v]++
    }
    minVal, minIdx := cnt[0], 0
    for i := 0; i < value; i++ {
        if cnt[i] < minVal {
            minIdx = i
            minVal = cnt[i]
        }
    }
    return minIdx + minVal*value
}

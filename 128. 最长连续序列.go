package main

func longestConsecutive(nums []int) int {
    m := make(map[int]bool) // 记录每个数字是否出现过
    for _, x := range nums {
        m[x] = true
    }
    ans := 0
    for k := range m {
        if m[k-1] { // 如果 k-1 出现过，说明 k 不是连续序列的起点，跳过
            continue
        }
        tmp := 1
        for m[k+1] { // 从k+1开始找连续的数字
            tmp++
            k++
        }
        ans = max(ans, tmp)
    }
    return ans
}

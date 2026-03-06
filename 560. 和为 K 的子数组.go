package main

func subarraySum(nums []int, k int) int {
    ans := 0
    // 问题变为前缀和中sum[j] - sum[i] == k的 sum[i] 的数量。 即两数之和。 a + b = k
    sum := 0
    mp := make(map[int]int)
    mp[0] = 1 // 与两数之和不同：一个sum[j]==k也可以，所以mp[0]=1
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        ans += mp[sum-k] // 要找的是 sum-k的数量
        mp[sum]++
    }

    return ans
}

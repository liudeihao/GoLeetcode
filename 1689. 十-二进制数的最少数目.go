package main

func minPartitions(n string) int {
    var ans int
    for _, c := range n {
        ans = max(ans, int(c-'0'))
    }
    return ans
}

package main

func minOperations(s string, k int) int {
    n := len(s)
    ones := 0
    for _, c := range s {
        if c == '1' {
            ones++
        }
    }
    onesVisited := make([]int, n+1)
    for i := 0; i <= n; i++ {
        onesVisited[i] = -1
    }
    onesVisited[ones] = 0
    q := []int{ones}
    for len(q) > 0 {
        cq := q
        q = nil
        for _, one := range cq {
            zero := n - one
            for i := max(0, k-one); i <= min(zero, k); i++ {
                next := one - k + 2*i
                if next >= 0 && next <= n && onesVisited[next] == -1 {
                    onesVisited[next] = onesVisited[one] + 1
                    q = append(q, next)
                }
            }
        }
    }
    return onesVisited[n]
}

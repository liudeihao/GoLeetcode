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

    parent := [2][]int{
        make([]int, n/2+2),
        make([]int, n/2+2),
    }
    for i := 0; i <= n/2+1; i++ {
        parent[0][i] = i
        parent[1][i] = i
    }
    // 路径压缩的并查集，奇偶分开。专门记录是否访问过。
    var find func(p, x int) int
    find = func(p, x int) int {
        if parent[p][x] != x {
            parent[p][x] = find(p, parent[p][x])
        }
        return parent[p][x]
    }

    // 初始状态标记为已访问。
    p := ones % 2
    parent[p][ones/2] = find(p, ones/2+1)

    q := []int{ones}
    for len(q) > 0 {
        one := q[0]
        q = q[1:]
        zero := n - one

        // I的取值范围
        minI := max(0, k-one)
        maxI := min(zero, k)

        // 下一步可能到达的区间 [L,R]
        L := one - k + 2*minI
        R := one - k + 2*maxI

        // 区间的奇偶性
        pNext := L % 2

        for {
            idx := find(pNext, L/2)
            nextVal := idx*2 + pNext
            if nextVal > R {
                break
            }
            onesVisited[nextVal] = onesVisited[one] + 1
            if nextVal == n {
                return onesVisited[n]
            }
            q = append(q, nextVal)
            // 核心！将当前节点标记为已访问，直接连到它的下一个节点。
            parent[pNext][idx] = find(pNext, idx+1)
        }
    }
    return onesVisited[n]
}

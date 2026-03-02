package main

func minSwaps(grid [][]int) int {
    n := len(grid)
    pos := make([]int, n)
    for i := 0; i < n; i++ {
        for j := n - 1; j >= 0; j-- {
            if grid[i][j] == 1 {
                pos[i] = j
                break
            }
        }
    }

    var ans int
    for i := 0; i < n; i++ {
        if pos[i] <= i {
            continue
        }
        k := -1
        for j := i + 1; j < n; j++ {
            if pos[j] <= i {
                ans += j - i
                k = j
                break
            }
        }
        if k == -1 {
            return -1
        } else {
            copy(pos[i+1:k+1], pos[i:k])
        }
    }
    return ans
}

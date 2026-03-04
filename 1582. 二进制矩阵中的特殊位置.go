package main

func numSpecial(mat [][]int) int {
    rows, cols := len(mat), len(mat[0])
    rowSum := make([]int, rows)
    colSum := make([]int, cols)
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            rowSum[i] += mat[i][j]
            colSum[j] += mat[i][j]
        }
    }

    var ans int
    for i := 0; i < rows; i++ {
        if rowSum[i] != 1 {
            continue
        }
        for j := 0; j < cols; j++ {
            if mat[i][j] == 1 && colSum[j] == 1 {
                ans++
            }
        }
    }
    return ans
}

package main

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	firstCol, firstRow := -1, -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				if firstRow == -1 && firstCol == -1 {
					firstRow = i
					firstCol = j
				}
				// matrix[firstRow][:]记录哪些列需要清空
				matrix[firstRow][j] = 0
				// matrix[:][firstCol]记录哪些行需要清空
				matrix[i][firstCol] = 0
			}
		}
	}

	if firstRow == -1 && firstCol == -1 {
		return
	}

	for i := 0; i < m; i++ {
		if i == firstRow {
			continue
		}
		for j := 0; j < n; j++ {
			if j == firstCol {
				continue
			}
			if matrix[firstRow][j] == 0 || matrix[i][firstCol] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	for i := 0; i < m; i++ {
		matrix[i][firstCol] = 0
	}

	for j := 0; j < n; j++ {
		matrix[firstRow][j] = 0
	}
}

/*
// 也可以直接用第一行第一列记录

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	// 用矩阵第一行第一列代替额外需要的数组
	// 为此需要记录第一行第一列是否会变为0
	row0, col0 := false, false
	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
			break
		}
	}
	for _, v := range matrix {
		if v[0] == 0 {
			col0 = true
			break
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				// 只是用矩阵第一行第一列代替额外需要的数组而已。
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row0 {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if col0 {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

*/

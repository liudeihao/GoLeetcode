package main

func solveSudoku(board [][]byte) {
	col := [10][10]bool{}
	row := [10][10]bool{}
	grid := [3][3][10]bool{}
	type pair struct {
		x, y int
	}
	var spaces []pair
	for i := range board {
		for j := range board {
			if board[i][j] == '.' {
				spaces = append(spaces, pair{i, j})
			} else {
				k := board[i][j] - '0'
				row[i][k] = true
				col[j][k] = true
				grid[i/3][j/3][k] = true
			}
		}
	}

	var backtrack func(int) bool
	backtrack = func(idx int) bool {
		if idx == len(spaces) {
			return true
		}
		p := spaces[idx]
		x, y := p.x, p.y
		for k := 1; k <= 9; k++ {
			if row[x][k] || col[y][k] || grid[x/3][y/3][k] {
				continue
			}
			row[x][k] = true
			col[y][k] = true
			grid[x/3][y/3][k] = true

			board[x][y] = byte('0' + k)
			if backtrack(idx + 1) {
				return true
			}
			board[x][y] = '.'

			row[x][k] = false
			col[y][k] = false
			grid[x/3][y/3][k] = false
		}
		return false
	}
	backtrack(0)
}

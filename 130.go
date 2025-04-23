package main

func solve(board [][]byte) {
	m, n := len(board), len(board[0])

	var dfs func(int, int)
	dx, dy := []int{-1, 0, 0, 1}, []int{0, -1, 1, 0}
	dfs = func(x, y int) {
		board[x][y] = '-'

		for k := range 4 {
			i, j := x+dx[k], y+dy[k]
			if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
				continue
			}
			dfs(i, j)
		}
	}
	for i := range m {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][n-1] == 'O' {
			dfs(i, n-1)
		}
	}
	for j := range n {
		if board[0][j] == 'O' {
			dfs(0, j)
		}
		if board[m-1][j] == 'O' {
			dfs(m-1, j)
		}
	}
	for i := range m {
		for j := range n {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '-' {
				board[i][j] = 'O'
			}
		}
	}

}

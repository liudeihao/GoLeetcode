package main

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	nb := make([][]int, m)
	for i := range m {
		nb[i] = make([]int, n)
	}

	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	countNeighbors := func(i, j int) (cnt int) {
		for k := range 8 {
			x, y := i+dx[k], j+dy[k]
			if x < 0 || y < 0 || x >= m || y >= n {
				continue
			} else if board[x][y] == 1 {
				cnt++
			}
		}
		return
	}
	for i := range m {
		for j := range n {
			neighbors := countNeighbors(i, j)
			if neighbors == 3 {
				nb[i][j] = 1
			} else if neighbors < 2 || neighbors > 3 {
				nb[i][j] = 0
			} else {
				nb[i][j] = board[i][j]
			}
		}
	}
	copy(board, nb)
}

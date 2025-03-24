package main

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range m {
		visited[i] = make([]bool, n)
	}
	var recurFunc func(string, int, int) bool
	recurFunc = func(w string, x, y int) (ans bool) {
		if len(w) == 0 {
			return true
		}
		if x < 0 || x >= m || y < 0 || y >= n || visited[x][y] || board[x][y] != w[0] {
			return
		}
		w = w[1:]
		visited[x][y] = true
		ans = ans || recurFunc(w, x-1, y)
		ans = ans || recurFunc(w, x+1, y)
		ans = ans || recurFunc(w, x, y-1)
		ans = ans || recurFunc(w, x, y+1)
		visited[x][y] = false
		return
	}
	for i := range m {
		for j := range n {
			if recurFunc(word, i, j) {
				return true
			}
		}
	}
	return false
}

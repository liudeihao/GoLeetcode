package main

func solveNQueens(n int) [][]string {
	type pair struct {
		x, y int
	}
	var ans [][]string
	var path []int
	ansAppend := func() {
		template := make([][]byte, n)
		for i := range template {
			template[i] = make([]byte, n)
			for j := range n {
				template[i][j] = '.'
			}
		}
		for i, j := range path {
			template[i][j] = 'Q'
		}
		tmp := make([]string, n)
		for i, line := range template {
			tmp[i] = string(line)
		}
		ans = append(ans, tmp)
	}

	// n*n的棋盘，n个竖线，2*n-1个对角线/斜对角线
	cols := make([]bool, n+1)
	diags := make([]bool, 2*n)
	idiags := make([]bool, 2*n)

	var backtrack func(int)
	backtrack = func(i int) {
		if len(path) == n {
			ansAppend()
			return
		}
		for j := 0; j < n; j++ {
			d, id := i+j, i-j+n
			if cols[j] || diags[d] || idiags[id] {
				// 所以后面设置为false不会影响到之前的queens
				continue
			}
			cols[j] = true
			diags[d] = true
			idiags[id] = true

			path = append(path, j)
			backtrack(i + 1)
			path = path[:len(path)-1]

			cols[j] = false
			diags[d] = false
			idiags[id] = false
		}
	}
	backtrack(0)
	return ans
}

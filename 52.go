package main

func totalNQueens(n int) (ans int) {
	type pos struct {
		x, y int
	}
	queens := make([]pos, n)

	collide := func(p pos) bool {
		for _, q := range queens {
			// vertical
			if q.y == p.y {
				return true
			}
			// diagonal
			if p.x-q.x == p.y-q.y || p.x-q.x == q.y-p.y {
				return true
			}
		}
		return false
	}

	var backtrace func()
	backtrace = func() {
		if len(queens) == n {
			ans++
			return
		}

		for j := 0; j < n; j++ {
			np := pos{len(queens), j}
			if collide(np) {
				continue
			}
			queens = append(queens, np)
			backtrace()
			queens = queens[:len(queens)-1]
		}
	}
	for j := 0; j < n; j++ {
		queens = nil
		queens = append(queens, pos{0, j})
		backtrace()
	}
	return
}

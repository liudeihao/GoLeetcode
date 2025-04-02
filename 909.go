package main

import (
	"slices"
)

func snakesAndLadders(board [][]int) int {
	n := len(board)
	for i := n - 2; i >= 0; i -= 2 {
		slices.Reverse(board[i])
	}
	arr := make([]int, 1)
	for i := n - 1; i >= 0; i-- {
		arr = append(arr, board[i]...)
	}
	vs := make([]int, n*n+1)

	q := []int{1}
	step := 1
	for len(q) > 0 {
		m := len(q)
		for _ = range m {
			front := q[0]
			q = q[1:]
			for j := 1; front+j <= n*n && j <= 6; j++ {
				next := front + j
				if arr[next] != -1 {
					next = arr[next]
				}

				if next == n*n {
					return step
				}

				if vs[next] != 0 {
					continue
				}
				vs[next] = step
				q = append(q, next)
			}
		}
		step++
	}
	return -1
}

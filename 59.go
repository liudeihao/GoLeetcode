package main

func generateMatrix(n int) [][]int {
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
	}
	p := 1
	l, r, t, b := 0, n-1, 0, n-1
	for {
		for i := l; i <= r; i++ {
			mat[t][i] = p
			p++
		}
		t++
		if t > b {
			break
		}
		for i := t; i <= b; i++ {
			mat[i][r] = p
			p++
		}
		r--
		if r < l {
			break
		}
		for i := r; i >= l; i-- {
			mat[b][i] = p
			p++
		}
		b--
		if b < t {
			break
		}
		for i := b; i >= t; i-- {
			mat[i][l] = p
			p++
		}
		l++
		if l > r {
			break
		}
	}
	return mat
}

package main

func spiralOrder(matrix [][]int) []int {
	var ans []int
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	for {
		for i := l; i <= r; i++ {
			ans = append(ans, matrix[t][i])
		}
		t++
		if t > b {
			break // 随时可能break
		}
		for i := t; i <= b; i++ {
			ans = append(ans, matrix[i][r])
		}
		r--
		if r < l {
			break
		}
		for i := r; i >= l; i-- {
			ans = append(ans, matrix[b][i])
		}
		b--
		if b < t {
			break
		}
		for i := b; i >= t; i-- {
			ans = append(ans, matrix[i][l])
		}
		l++
		if l > r {
			break
		}
	}
	return ans
}

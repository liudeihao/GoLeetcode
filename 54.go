package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	if m == 1 {
		return matrix[0]
	} else if n == 1 {
		ans := make([]int, m)
		for i := range m {
			ans[i] = matrix[i][0]
		}
		return ans
	}

	ans := make([]int, (n+m-2)*2)
	p := 0

	for j := range n - 1 {
		ans[p] = matrix[0][j]
		p++
	}
	for i := range m - 1 {
		ans[p] = matrix[i][n-1]
		p++
	}
	for j := range n - 1 {
		ans[p] = matrix[m-1][n-1-j]
		p++
	}
	for i := range m - 1 {
		ans[p] = matrix[m-1-i][0]
		p++
	}

	innerMatrix := make([][]int, m-2)
	for i := range innerMatrix {
		innerMatrix[i] = make([]int, n-2)
		for j := range innerMatrix[i] {
			innerMatrix[i][j] = matrix[i+1][j+1]
		}
	}
	return append(ans, spiralOrder(innerMatrix)...)
}

func main() {
	for i := range 1 {
		fmt.Println(i)
	}
}

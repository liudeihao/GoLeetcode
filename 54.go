package main

import "fmt"

/*
func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return []int{}
    }
    rows, columns := len(matrix), len(matrix[0])
    visited := make([][]bool, rows)
    for i := 0; i < rows; i++ {
        visited[i] = make([]bool, columns)
    }

    var (
        total = rows * columns
        order = make([]int, total)
        row, column = 0, 0
        directions = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
        directionIndex = 0
    )

    for i := 0; i < total; i++ {
        order[i] = matrix[row][column]
        visited[row][column] = true
        nextRow, nextColumn := row + directions[directionIndex][0], column + directions[directionIndex][1]
        if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
            directionIndex = (directionIndex + 1) % 4
        }
        row += directions[directionIndex][0]
        column += directions[directionIndex][1]
    }
    return order
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/spiral-matrix/solutions/275393/luo-xuan-ju-zhen-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func spiralOrder(matrix [][]int) []int {
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	var ans []int
	for {
		for i := l; i <= r; i++ {
			ans = append(ans, matrix[t][i])
		}
		t++
		if t > b {
			break
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

func main() {
	for i := range 1 {
		fmt.Println(i)
	}
}

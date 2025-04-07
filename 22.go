package main

import "fmt"

func generateParenthesis(n int) []string {
	var ans []string
	var path []byte
	l, r := 0, 0
	var backtrack func()
	backtrack = func() {
		if r > l || l > n {
			return
		}
		if r == l && len(path) == n*2 {
			ans = append(ans, string(path))
			return
		}
		l++
		path = append(path, '(')
		backtrack()
		l--
		r++
		path[len(path)-1] = ')'
		backtrack()
		r--
		path = path[:len(path)-1]
	}
	backtrack()
	return ans
}

func main() {
	fmt.Println(generateParenthesis(3))
}

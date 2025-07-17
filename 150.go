package main

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, num)
		} else {
			size := len(stack)
			var res int
			n1, n2 := stack[size-2], stack[size-1]
			stack = stack[:size-2]
			switch token {
			case "+":
				res = n1 + n2
			case "-":
				res = n1 - n2
			case "*":
				res = n1 * n2
			case "/":
				res = n1 / n2
			}
			stack = append(stack, res)
		}
	}
	return stack[0]
}

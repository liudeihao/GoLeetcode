package main

func isValid(s string) bool {
	var stack []rune
	for _, c := range s {
		// 有一些技巧，在匹配左括号的时候，右括号先入栈，
		// 就只需要比较当前元素和栈顶相不相等就可以了，比左括号先入栈代码实现要简单的多了！
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
			continue
		} else if len(stack) == 0 {
			return false
		}
		top := stack[len(stack)-1]
		if c == ')' && top != '(' || c == ']' && top != '[' || c == '}' && top != '{' {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}

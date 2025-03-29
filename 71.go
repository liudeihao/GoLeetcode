package main

import "strings"

func splitNames(path string) (names []string) {
	n := len(path)
	l, r := 1, 1
	for r < n {
		for l < n && path[l] == '/' {
			l++
			r++
		}
		for r < n && path[r] != '/' {
			r++
		}
		if l < r {
			names = append(names, path[l:r])
		}
		l = r
	}
	return
}

func simplifyPath(path string) string {
	names := splitNames(path)
	var stack []string
	for _, name := range names {
		if name == ".." && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else if name != ".." && name != "." {
			stack = append(stack, name)
		}
	}
	return "/" + strings.Join(stack, "/")
}

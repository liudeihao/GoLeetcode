package main

import "strings"

func wordPattern(pattern string, s string) bool {
	ss := strings.Split(s, " ")
	c2s := make(map[int32]string)
	s2c := make(map[string]int32)
	if len(pattern) != len(ss) {
		return false
	}
	for i, c := range pattern {
		s := ss[i]
		if c2s[c] == "" {
			c2s[c] = s
		}
		if s2c[s] == 0 {
			s2c[s] = c
		}
		if c2s[c] != s || s2c[s] != c {
			return false
		}
	}
	return true
}

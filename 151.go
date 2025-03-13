package main

import (
	"bytes"
	"slices"
)

func splitWords(s string) [][]byte {
	n := len(s)
	var words [][]byte
	for start, end := 0, 0; end < n; {
		for start < n && s[start] == ' ' {
			start++
		}
		for end < n && s[end] == ' ' {
			end++
		}
		for end < n && s[end] != ' ' {
			end++
		}
		if start < end {
			words = append(words, []byte(s[start:end]))
			start = end
		}
	}
	return words
}

func reverseWords(s string) string {
	words := splitWords(s)
	slices.Reverse(words)
	return string(bytes.Join(words, []byte(" ")))
}

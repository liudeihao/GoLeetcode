package main

import (
	"bytes"
	"slices"
)

func removeAnagrams(words []string) []string {
	var last []byte
	l := 0
	for _, word := range words {
		s := []byte(word)
		slices.Sort(s)
		if bytes.Compare(s, last) != 0 {
			last = s
			words[l] = word
			l++
		}
	}
	return words[:l]
}

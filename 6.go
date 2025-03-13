package main

import "bytes"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	buckets := make([][]byte, numRows)
	i, curRow := 0, 0
	n := len(s)
	for i < n {
		for i < n && curRow < numRows-1 {
			buckets[curRow] = append(buckets[curRow], s[i])
			curRow++
			i++
		}
		for i < n && curRow > 0 {
			buckets[curRow] = append(buckets[curRow], s[i])
			curRow--
			i++
		}
	}
	return string(bytes.Join(buckets, nil))
}

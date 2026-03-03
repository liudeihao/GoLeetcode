package main

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(turnedOn int) []string {
	var ans []string
	hours := [4][]int{
		{0},
		{1, 2, 4, 8},
		{3, 5, 6, 9, 10},
		{7, 11},
	}
	var minutes [60]int
	for i := 0; i < 60; i++ {
		minutes[i] = bits.OnesCount8(uint8(i))
	}

	for i := 0; i <= turnedOn && i < 4; i++ {
		var ms []int
		for j := 0; j < 60; j++ {
			if minutes[j]+i == turnedOn {
				ms = append(ms, j)
			}
		}
		for _, h := range hours[i] {
			for _, m := range ms {
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return ans
}

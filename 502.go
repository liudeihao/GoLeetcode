package main

import (
	"container/heap"
	"sort"
)

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *hp) Push(v any) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() any {
	a := h.IntSlice
	n := len(h.IntSlice)
	v := a[n-1]
	h.IntSlice = a[:n-1]
	return v
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	type pair struct{ c, p int }
	arr := make([]pair, n)
	for i, p := range profits {
		arr[i] = pair{capital[i], p}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].c < arr[j].c
	})

	h := &hp{}
	for cur := 0; k > 0; k-- {
		for cur < n && arr[cur].c <= w {
			heap.Push(h, arr[cur].p)
			cur++
		}
		if h.Len() == 0 {
			break
		}
		w += heap.Pop(h).(int)
	}
	return w
}

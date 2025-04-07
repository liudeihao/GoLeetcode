package main

import (
	"container/heap"
)

type Pair struct {
	x, y int
}

type PairSlice []Pair

func (ps PairSlice) Len() int {
	return len(ps)
}

func (ps PairSlice) Less(i, j int) bool {
	p1, p2 := ps[i], ps[j]
	return p1.x+p1.y > p2.x+p2.y
}

func (ps PairSlice) Swap(i, j int) {
	p1, p2 := ps[i], ps[j]
	tmp := p1
	p1 = p2
	p2 = tmp
}

func (ps *PairSlice) Push(x any) {
	*ps = append(*ps, x.(Pair))
}

func (ps *PairSlice) Pop() any {
	n := len(*ps)
	tail := (*ps)[n-1]
	*ps = (*ps)[:n-1]
	return tail
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n := len(nums1)
	ps := make(PairSlice, n)
	for i := range n {
		ps[i] = Pair{nums1[i], nums2[i]}
	}
	heap.Init(&ps)
	var ans [][]int
	for _ = range k {
		top := heap.Pop(&ps).(Pair)
		ans = append(ans, []int{top.x, top.y})
	}
	return ans
}

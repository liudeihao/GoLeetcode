package main

import "container/heap"

type Pair struct {
	i, j int
	x, y int
}

type PairSlice struct {
	pairs []Pair
}

func (ps PairSlice) Len() int {
	return len(ps.pairs)
}

func (ps PairSlice) Less(i, j int) bool {
	p1, p2 := ps.pairs[i], ps.pairs[j]
	return p1.x+p1.y < p2.x+p2.y
}

func (ps PairSlice) Swap(i, j int) {
	tmp := ps.pairs[i]
	ps.pairs[i] = ps.pairs[j]
	ps.pairs[j] = tmp
}

func (ps *PairSlice) Push(x any) {
	ps.pairs = append(ps.pairs, x.(Pair))
}

func (ps *PairSlice) Pop() any {
	a := ps.pairs
	n := ps.Len()
	last := a[n-1]
	ps.pairs = a[:n-1]
	return last
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n1, n2 := len(nums1), len(nums2)
	ps := &PairSlice{make([]Pair, n1)}

	for i := range n1 {
		ps.pairs[i] = Pair{i, 0, nums1[i], nums2[0]}
	}
	var ans [][]int
	heap.Init(ps)
	for _ = range k {
		top := heap.Pop(ps).(Pair)
		ans = append(ans, []int{top.x, top.y})
		if top.j+1 < n2 {
			heap.Push(ps, Pair{top.i, top.j + 1, nums1[top.i], nums2[top.j+1]})
		}
	}
	return ans
}

package main

import "container/heap"

type Pair struct {
	i, j int // 只需要存索引！
}

type PairSlice struct { // 其实命名hp， 然后func (h hp) ... 合适吧
	nums1, nums2 []int // 就可以通过索引获取数据哩！
	data         []Pair
}

func (ps PairSlice) Len() int {
	return len(ps.data)
}

func (ps PairSlice) Less(i, j int) bool {
	p1, p2 := ps.data[i], ps.data[j]
	return ps.nums1[p1.i]+ps.nums2[p1.j] < ps.nums1[p2.i]+ps.nums2[p2.j]
}

func (ps PairSlice) Swap(i, j int) {
	// go语言是支持这个的...
	ps.data[i], ps.data[j] = ps.data[j], ps.data[i]
}

func (ps *PairSlice) Push(v any) {
	ps.data = append(ps.data, v.(Pair))
}

func (ps *PairSlice) Pop() any {
	a := ps.data
	n := ps.Len()
	last := a[n-1]
	ps.data = a[:n-1]
	return last
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n1, n2 := len(nums1), len(nums2)
	ps := &PairSlice{nums1, nums2, make([]Pair, n1)}

	for i := range n1 {
		ps.data[i] = Pair{i, 0}
	}

	var ans [][]int
	heap.Init(ps)
	for _ = range k {
		top := heap.Pop(ps).(Pair)
		i, j := top.i, top.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n2 {
			heap.Push(ps, Pair{i, j + 1})
		}
	}
	return ans
}

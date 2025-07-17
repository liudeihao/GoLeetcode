package main

import (
	"container/heap"
)

type pair struct {
	num, cnt int
}

type PriorityQueue struct {
	data []pair
}

func (pq *PriorityQueue) Len() int {
	return len(pq.data)
}
func (pq *PriorityQueue) Swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.data[i].cnt > pq.data[j].cnt
}

func (pq *PriorityQueue) Push(x any) {
	pq.data = append(pq.data, x.(pair))
}

func (pq *PriorityQueue) Pop() any {
	n := len(pq.data)
	top := pq.data[n-1]
	pq.data = pq.data[0 : n-1]
	return top
}

func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}

	pq := &PriorityQueue{}
	for key, val := range mp {
		heap.Push(pq, pair{key, val})
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}

	var ans []int
	for i := 0; i < k; i++ {
		p := pq.Pop()
		ans = append(ans, p.(pair).num)
	}
	return ans
}

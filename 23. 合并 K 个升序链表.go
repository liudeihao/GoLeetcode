package main

import "container/heap"

type lheap []*ListNode

func (l lheap) Len() int {
	return len(l)
}

func (l lheap) Less(i, j int) bool {
	return l[i].Val < l[j].Val
}
func (l lheap) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l *lheap) Push(x any) {
	*l = append(*l, x.(*ListNode))
}

func (l *lheap) Pop() any {
	last := (*l)[l.Len()-1]
	*l = (*l)[:l.Len()-1]
	return last
}

func mergeKLists(lists []*ListNode) *ListNode {
	// 把这些垃圾都清理掉先
	i := 0
	for j := 0; j < len(lists); j++ {
		if lists[j] != nil {
			lists[i] = lists[j]
			i++
		}
	}
	lists = lists[:i]

	hp := lheap(lists)
	heap.Init(&hp) // 记得init
	dummy := &ListNode{}
	p := dummy
	for hp.Len() > 0 {
		top := heap.Pop(&hp).(*ListNode)
		p.Next = top
		p = p.Next
		if top.Next != nil {
			heap.Push(&hp, top.Next)
		}
	}
	return dummy.Next
}

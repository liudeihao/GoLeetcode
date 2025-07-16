package main

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{-1, head}
	p := dummy

	for p.Next != nil && p.Next.Next != nil {
		next1 := p.Next
		next2 := p.Next.Next
		next3 := p.Next.Next.Next
		next1.Next = next3
		next2.Next = next1
		p.Next = next2
		p = next1
	}
	return dummy.Next
}

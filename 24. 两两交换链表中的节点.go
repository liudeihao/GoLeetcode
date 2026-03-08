package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	nx := p.Next
	nnx := p.Next.Next
	nx.Next = p
	p.Next = swapPairs(nnx)

	return nx
}

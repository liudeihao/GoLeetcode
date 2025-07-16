package main

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p, pre := head.Next, head
	pre.Next = nil // 这个不能忘啊
	for p != nil {
		next := p.Next
		p.Next = pre
		pre = p
		p = next
	}
	return pre
}

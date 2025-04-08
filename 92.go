package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	n := right - left
	dummy := &ListNode{
		-1, head,
	}
	lpre := dummy
	for ; left > 1; left-- {
		lpre = lpre.Next
	}
	l := lpre.Next
	next := l.Next
	l.Next = nil
	for ; n > 0; n-- {
		pre := l
		l = next
		next = l.Next
		l.Next = pre
	}
	lpre.Next.Next = next
	lpre.Next = l
	return dummy.Next
}

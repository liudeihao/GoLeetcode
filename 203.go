package main

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	for p := dummy; p.Next != nil; {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return dummy.Next
}

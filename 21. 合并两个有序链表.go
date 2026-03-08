package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	dummy := &ListNode{}
	p := dummy
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			next := p1.Next
			p.Next = p1
			p = p.Next
			p1 = next
		} else {
			next := p2.Next
			p.Next = p2
			p = p.Next
			p2 = next
		}
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return dummy.Next
}

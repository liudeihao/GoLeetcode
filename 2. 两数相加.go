package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	carry := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		digit := sum % 10
		carry = sum / 10
		p.Next = &ListNode{Val: digit}
		p = p.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		sum := l1.Val + carry
		digit := sum % 10
		carry = sum / 10
		p.Next = &ListNode{Val: digit}
		p = p.Next
		l1 = l1.Next
	}
	for l2 != nil {
		sum := l2.Val + carry
		digit := sum % 10
		carry = sum / 10
		p.Next = &ListNode{Val: digit}
		p = p.Next
		l2 = l2.Next
	}
	if carry != 0 {
		p.Next = &ListNode{Val: 1}
	}
	return dummy.Next
}

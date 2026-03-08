package main

func isPalindrome(head *ListNode) bool {
	reverse := func(n *ListNode) *ListNode {
		// 反转链表
		p := n
		var prev *ListNode
		for p != nil {
			next := p.Next
			p.Next = prev
			prev = p
			p = next
		}
		return prev
	}
	middle := func(n *ListNode) *ListNode {
		// 当偶数个节点，返回第二个。
		slow, fast := head, head
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		return slow
	}
	m := middle(head)
	r := reverse(m)
	l := head
	for r != nil { // 偶数个节点时，会是这样的：1->2->3<-4
		if l.Val != r.Val {
			return false
		}
		l = l.Next
		r = r.Next
	}
	return true
}

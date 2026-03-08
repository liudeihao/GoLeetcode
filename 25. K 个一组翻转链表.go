package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil {
		return head
	}

	length := 0
	for q := head; q != nil && length < k; q = q.Next {
		length++ // 取K个节点
	}

	if length < k { // 剩余节点不足K个，不反转
		return head
	}

	curHead := head
	p := head
	var prev, next *ListNode

	for i := k; p != nil && i > 0; i-- { // 反转从p开始的K个节点
		next = p.Next
		p.Next = prev
		prev = p
		p = next
	}
	curHead.Next = reverseKGroup(next, k)
	return prev
}

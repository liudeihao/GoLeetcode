package main

func middleNode(head *ListNode) *ListNode {
	// 当中间有两个节点时，返回第二个中间结点。
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

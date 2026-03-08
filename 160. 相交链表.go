package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 如果 headA = x + z, headB = y + z
	// 注意到 (x + z) + y == (y + z) + x, 两个都走一遍就会相遇在交点
	p1, p2 := headA, headB
	for p1 != p2 {
		// 每次循环，各向后走一步
		if p1 != nil {
			p1 = p1.Next
		} else { // 走到空节点时，去走另一个链表
			p1 = headB
		}
		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = headA
		}
	}
	return p1
}

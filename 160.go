package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	for p := headA; p != nil; p = p.Next {
		lenA++
	}
	for p := headB; p != nil; p = p.Next {
		lenB++
	}
	if lenA > lenB {
		// A短
		headA, headB = headB, headA
		lenA, lenB = lenB, lenA
	}
	slow, fast := headA, headB
	diff := lenB - lenA
	for diff > 0 {
		diff--
		fast = fast.Next
	}
	for slow != fast && fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	if fast == nil {
		return nil
	}
	return slow
}

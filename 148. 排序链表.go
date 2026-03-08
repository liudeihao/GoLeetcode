package main

// 但是依然是O(log n)空间复杂度（递归栈）
func middleListNode(head *ListNode) (*ListNode, *ListNode) {
	slow, fast := head, head
	var slowPrev *ListNode
	for fast != nil && fast.Next != nil {
		slowPrev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slowPrev, slow
}

func mergeListNodes(h1, h2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			p.Next = h1
			h1 = h1.Next
			p = p.Next
		} else {
			p.Next = h2
			h2 = h2.Next
			p = p.Next
		}
	}
	for h1 != nil {
		p.Next = h1
		h1 = h1.Next
		p = p.Next
	}
	for h2 != nil {
		p.Next = h2
		h2 = h2.Next
		p = p.Next
	}
	return dummy.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	middlePrev, middle := middleListNode(head)
	if middlePrev != nil {
		middlePrev.Next = nil
	}
	left, right := sortList(head), sortList(middle)
	return mergeListNodes(left, right)
}

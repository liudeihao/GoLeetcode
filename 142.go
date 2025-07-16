package main

// 头节点---x---环形入口节点---y---fast与slow相遇节点---z---环形入口节点
// 则有 x + y + n(y+z) == 2(x + y)
// 等式左边是fast走过的节点数，右侧是slow走过的节点数*2
// 于是x = n(y+z)-y = (n-1)(y+z) + z
// 这等价于，从头节点出发一个指针、从相遇节点发出一个指针，那么他们相遇的地方就是环形入口节点
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			p1, p2 := head, slow
			for p1 != p2 {
				p1 = p1.Next
				p2 = p2.Next
			}
			return p1
		}
	}
	return nil
}

package main

type Node_117 struct {
	Val   int
	Left  *Node_117
	Right *Node_117
	Next  *Node_117
}

func connect_117(root *Node_117) *Node_117 {
	if root == nil {
		return root
	}
	q := []*Node_117{root}
	for len(q) > 0 {
		// 这个写法好好啊，我要学会
		tmp := q
		q = nil
		for i, node := range tmp {
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return root
}

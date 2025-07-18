package main

type Node_116 struct {
	Val   int
	Left  *Node_116
	Right *Node_116
	Next  *Node_116
}

func connect(root *Node_116) *Node_116 {
	if root == nil {
		return nil
	}
	var q []*Node_116
	q = append(q, root)
	for len(q) != 0 {
		cq := q
		q = nil
		for i, tn := range cq {
			if i+1 < len(cq) {
				tn.Next = cq[i+1]
			}
			if tn.Left != nil {
				q = append(q, tn.Left)
			}
			if tn.Right != nil {
				q = append(q, tn.Right)
			}
		}
	}
	return root
}

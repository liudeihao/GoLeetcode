package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tn *TreeNode) Show() {
	if tn == nil {
		fmt.Println("nil")
		return
	}
	tn.layerOrd(func(root *TreeNode) {
		fmt.Printf(" %d ", root.Val)
	})
}

func (tn *TreeNode) layerOrd(f func(*TreeNode)) {
	q := []*TreeNode{tn}
	for len(q) > 0 {
		cq := q
		q = nil
		for _, node := range cq {
			f(node)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}

		}
	}
}

func (tn *TreeNode) prevOrd(f func(*TreeNode)) {
	f(tn)
	if tn.Left != nil {
		tn.Left.prevOrd(f)
	}
	if tn.Right != nil {
		tn.Right.prevOrd(f)
	}
}

func (tn *TreeNode) midOrd(f func(*TreeNode)) {
	if tn.Left != nil {
		tn.Left.midOrd(f)
	}
	f(tn)
	if tn.Right != nil {
		tn.Right.midOrd(f)
	}
}

func (tn *TreeNode) postOrd(f func(*TreeNode)) {
	if tn.Left != nil {
		tn.Left.postOrd(f)
	}
	if tn.Right != nil {
		tn.Right.postOrd(f)
	}
	f(tn)

}

func NewTree(vals []int) *TreeNode {
	n := len(vals)
	if n == 0 {
		return nil
	}
	list := make([]*TreeNode, n)
	for i := range list {
		if vals[i] != 0 {
			list[i] = &TreeNode{Val: vals[i]}
		} else {
			list[i] = nil
		}
	}
	for i := range list {
		if list[i] == nil {
			continue
		}
		if i*2+1 < n {
			list[i].Left = list[i*2+1]
		}
		if i*2+2 < n {
			list[i].Right = list[i*2+2]
		}
	}
	return list[0]
}

package main

type BSTIterator struct {
	p  int
	ns []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	i := BSTIterator{}
	var midorder func(r *TreeNode)
	midorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		midorder(r.Left)
		i.ns = append(i.ns, r)
		midorder(r.Right)
	}
	midorder(root)
	return i
}

func (i *BSTIterator) Next() int {
	val := i.ns[i.p].Val
	i.p++
	return val
}

func (i *BSTIterator) HasNext() bool {
	return i.p != len(i.ns)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

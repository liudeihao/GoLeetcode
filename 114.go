package main

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var list []*TreeNode
	var prevOrd func(*TreeNode)
	prevOrd = func(node *TreeNode) {
		list = append(list, node)
		if node.Left != nil {
			prevOrd(node.Left)
		}
		if node.Right != nil {
			prevOrd(node.Right)
		}
	}
	prevOrd(root)
	for i := 0; i < len(list)-1; i++ {
		list[i].Left = nil
		list[i].Right = list[i+1]
	}
}

package main

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}

	var insertIntoBSTHelper func(*TreeNode, int)
	insertIntoBSTHelper = func(root *TreeNode, val int) {
		if val < root.Val {

			if root.Left == nil {
				root.Left = &TreeNode{val, nil, nil}
			} else {
				insertIntoBSTHelper(root.Left, val)
			}
		} else {
			
			if root.Right == nil {
				root.Right = &TreeNode{val, nil, nil}
			} else {
				insertIntoBSTHelper(root.Right, val)
			}
		}
	}

	insertIntoBSTHelper(root, val)
	return root
}

package main

func buildTree_105(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return &TreeNode{preorder[0], nil, nil}
	}
	t := preorder[0]
	llen := 0
	for i := range inorder {
		if inorder[i] == t {
			llen = i
			break
		}
	}
	return &TreeNode{
		preorder[0],
		buildTree_105(preorder[1:llen+1], inorder[:llen]),
		buildTree_105(preorder[llen+1:], inorder[llen+1:]),
	}
}

package main

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}
	rootVal := postorder[n-1]
	var inorderIdx int
	for i, val := range inorder {
		if val == rootVal {
			inorderIdx = i
			break
		}
	}
	llen, rlen := inorderIdx, n-inorderIdx-1
	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(inorder[:llen], postorder[:llen]),
		Right: buildTree(inorder[n-rlen:], postorder[n-rlen-1:]),
	}
}

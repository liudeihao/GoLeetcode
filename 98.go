package main

func isValidBST(root *TreeNode) bool {
	var res []int
	var trav func(node *TreeNode)
	trav = func(node *TreeNode) {
		if node == nil {
			return
		}
		trav(node.Left)
		res = append(res, node.Val)
		trav(node.Right)
	}
	trav(root)
	for i := 1; i < len(res); i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}
	return true
}

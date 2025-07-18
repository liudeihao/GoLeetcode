package main

func pathSum(root *TreeNode, targetSum int) [][]int {
	var ans [][]int
	var path []int
	var recursion func(root *TreeNode, targetSum int)
	recursion = func(root *TreeNode, targetSum int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		if targetSum == root.Val && root.Left == nil && root.Right == nil {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
		}
		recursion(root.Left, targetSum-root.Val)
		recursion(root.Right, targetSum-root.Val)
		path = path[:len(path)-1]
	}
	recursion(root, targetSum)
	return ans
}

package main

func sumRootToLeaf(root *TreeNode) int {
    var dfs func(*TreeNode, int) int
    dfs = func(root *TreeNode, curval int) int {
        if root == nil {
            return 0
        }
        curval = curval<<1 + root.Val
        if root.Left == nil && root.Right == nil {
            return curval
        }
        return dfs(root.Left, curval) + dfs(root.Right, curval)
    }
    return dfs(root, 0)
}
